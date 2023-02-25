package pilabitacora

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
	next  *Nodo
	value *est.Estudiante
	fecha string
	hora  string
}

type Pila struct {
	cabeza *Nodo
}

// Constructores

func NuevaPila() *Pila {
	p := new(Pila)
	p.cabeza = nil
	return p
}

// Insertar datos

func (p *Pila) Push(e *est.Estudiante) {
	now := time.Now()
	nodo := &Nodo{
		value: e,
		fecha: now.Format(time.Kitchen),
		hora:  now.Format("2006-02-01"),
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
	fmt.Println("Archivo creado exitosamente", nombre_archivo)
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
	fmt.Println("Archivo actualizado existosamente.")
}

func ejecutar(nombre_imagen string, archivo_dot string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo_dot).Output()
	mode := 0777
	_ = ioutil.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
}

func (l *Pila) Graficar(nombreArchivo string) {
	fmt.Println("Impresion")
	nombre_archivo_dot := fmt.Sprintf("./%s.dot", nombreArchivo)
	nombre_imagen := fmt.Sprintf("%s.jpg", nombreArchivo)
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	auxiliar := l.cabeza
	contador := 0
	for i := 0; i < l.Size(); i++ {
		texto += fmt.Sprintf("nodo%s[label=\"{|%s\\n%s|}\"];\n", strconv.Itoa(i), auxiliar.value.GetCarnet(), auxiliar.value.GetNombre())
		//texto = texto + "nodo" + strconv.Itoa(i) + "[label=\"{|" + "valor: " + ", " + auxiliar.value.GetNombre() + "|}\"];\n"
		auxiliar = auxiliar.next
	}
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Size()-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		//texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"

	crearArchivoDot(nombre_archivo_dot)
	escribirArchivoDot(texto, nombre_archivo_dot)
	ejecutar(nombre_imagen, nombre_archivo_dot)
}
