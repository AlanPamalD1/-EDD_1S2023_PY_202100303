package pilabitacora

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Nodo struct {
	next  *Nodo
	fecha string
	hora  string
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

func (p *Pila) Push( /*e *est.Estudiante */ ) {
	now := time.Now()
	nodo := &Nodo{
		//value: e,
		hora:  now.Format(time.Kitchen),
		fecha: now.Format("2006-02-01"),
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

func (p *Pila) Print() {

	if p.cabeza == nil {
		fmt.Println("Pila vacía")
		return
	}

	actual := p.cabeza

	for actual != nil {
		//actual.value.Print()
		fmt.Printf("Hora: %s Fecha: %s\n", actual.hora, actual.fecha)
		actual = actual.next
		fmt.Printf("%s\n", strings.Repeat("-", 45))
	}
}

// Eleminar datos

func (p *Pila) Pop() /* *est.Estudiante */ {
	if p.cabeza == nil {
		return
	}
	//e := p.cabeza.value
	p.cabeza = p.cabeza.next
	//return e
}

func (p *Pila) Subgrafo(i int) string {
	texto := ""
	auxiliar := p.cabeza

	for j := 0; j < p.Size(); j++ {
		texto += fmt.Sprintf("	nodo%db%d[label=\"Se inició sesión en\\n%s %s\"];\n", i, j, auxiliar.fecha, auxiliar.hora)
		auxiliar = auxiliar.next
	}

	return texto
}

func (p *Pila) UnionSubgrafos(i int) string {
	if p.Size() > 0 {
		texto := fmt.Sprintf("	nodo%d -> nodo%db%d;\n", i, i, 0)

		for j := 0; j < p.Size()-1; j++ {
			texto += fmt.Sprintf("	nodo%db%d -> nodo%db%d;\n", i, j, i, (j + 1))
		}
		return texto
	}
	return ""
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
	fmt.Printf("Archivo %s creado exitosamente\n", nombre_imagen)
}

func (p *Pila) Graficar(nombreArchivo string) {
	nombre_archivo_dot := fmt.Sprintf("./%s.dot", nombreArchivo)
	nombre_imagen := fmt.Sprintf("%s.jpg", nombreArchivo)
	texto := "digraph lista{\n"
	texto += "	rankdir=TB;\n"
	texto += "	node[shape = rectangle];\n"
	texto += "	edge[arrowsize=0.5];\n"
	texto += "	ranksep=\"0.2 equally\";\n"
	texto += "	nodonull1[label=\"Entrada\"];\n"
	auxiliar := p.cabeza

	for i := 0; i < p.Size(); i++ {
		texto += fmt.Sprintf("	nodo%d[label=\"Se inició sesión en\\n%s %s\"];\n", i, auxiliar.fecha, auxiliar.hora)
		auxiliar = auxiliar.next
	}
	if p.Size() > 0 {
		texto += "	nodonull1 -> nodo0;\n"
		for i := 0; i < p.Size()-1; i++ {
			texto += fmt.Sprintf("	nodo%d -> nodo%d;\n", i, i+1)
		}
	}

	texto += "}"

	crearArchivoDot(nombre_archivo_dot)
	escribirArchivoDot(texto, nombre_archivo_dot)
	ejecutar(nombre_imagen, nombre_archivo_dot)
}
