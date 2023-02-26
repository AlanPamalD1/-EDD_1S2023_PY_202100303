package pilaregistro

import (
	est "estudiante"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Nodo struct {
	next       *Nodo
	value      *est.Estudiante
	fecha      string
	hora       string
	aceptacion string
}

type Pila struct {
	cabeza *Nodo
}

// Constructores

func New() *Pila {
	p := new(Pila)
	p.cabeza = nil
	return p
}

// Insertar datos

func (p *Pila) Push(e *est.Estudiante, estado bool) {
	now := time.Now()
	aceptacion := "rechazó"
	if estado {
		aceptacion = "aceptó"
	}
	nodo := &Nodo{
		value:      e,
		fecha:      now.Format("2006-02-01"),
		hora:       now.Format(time.Kitchen),
		aceptacion: aceptacion,
	}
	if p.cabeza == nil {
		p.cabeza = nodo
	} else {
		nodo.next = p.cabeza
		p.cabeza = nodo
	}
}

// Verificar datos

func (l *Pila) Size() int {
	actual := l.cabeza
	tamanio := 0
	for actual != nil {
		tamanio++
		actual = actual.next
	}
	return tamanio
}

func (l *Pila) Exist(carnet string) bool {
	nodoActual := l.cabeza
	for nodoActual != nil {
		if nodoActual.value.GetCarnet() == carnet {
			return true
		}
		nodoActual = nodoActual.next
	}
	return false
}

func (p *Pila) Print() {

	if p.cabeza == nil {
		fmt.Println("Pila vacía")
		return
	}

	actual := p.cabeza

	for actual != nil {
		actual.value.Print()
		actual = actual.next
		fmt.Printf("%s\n", strings.Repeat("-", 45))
	}
}

// Eleminar datos

func (p *Pila) Pop() *est.Estudiante {
	if p.cabeza == nil {
		return nil
	}
	e := p.cabeza.value
	p.cabeza = p.cabeza.next
	return e
}

// Creación archivos dot

func crearArchivoDot(nombre_archivo string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(nombre_archivo)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(nombre_archivo)
		if err != nil {
			return
		}
		defer file.Close()
	}
}

func escribirArchivoDot(contenido string, nombre_archivo string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(nombre_archivo, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	_, err = file.WriteString(contenido)
	if err != nil {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if err != nil {
		return
	}
}

func ejecutar(nombre_imagen string, archivo_dot string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo_dot).Output()
	mode := 0777
	_ = ioutil.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
	fmt.Printf("Archivo %s.jpg creado exitosamente\n", nombre_imagen)
}

func (l *Pila) Graficar(nombreArchivo string) {
	nombre_archivo_dot := fmt.Sprintf("./%s.dot", nombreArchivo)
	nombre_imagen := fmt.Sprintf("%s.jpg", nombreArchivo)
	texto := "digraph lista{\n"
	texto += "	rankdir=TB;\n"
	texto += "	node[shape = rectangle];\n"
	texto += "	ranksep=\"0.05 equally\";\n"
	auxiliar := l.cabeza
	for i := 0; i < l.Size(); i++ {
		texto += fmt.Sprintf("	nodo%s[label=\"Se %s a\\n%s\\n %s %s\"];\n", strconv.Itoa(i), auxiliar.aceptacion, auxiliar.value.GetNombre(), auxiliar.fecha, auxiliar.hora)
		auxiliar = auxiliar.next
	}
	for i := 0; i < l.Size()-1; i++ {
		texto += fmt.Sprintf("	nodo%d -> nodo%d;\n", i, i+1)
	}
	texto += "}"

	crearArchivoDot(nombre_archivo_dot)
	escribirArchivoDot(texto, nombre_archivo_dot)
	ejecutar(nombre_imagen, nombre_archivo_dot)
}
