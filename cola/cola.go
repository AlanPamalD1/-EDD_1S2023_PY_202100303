package cola

import (
	est "estudiante"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Nodo struct {
	next  *Nodo
	value *est.Estudiante
}

type Cola struct {
	cabeza *Nodo
	cola   *Nodo
}

// Constructores

func New() *Cola {
	l := new(Cola)
	l.cabeza = nil
	l.cola = nil
	return l
}

// Insertar datos

func (c *Cola) Add(value *est.Estudiante) {
	newNode := &Nodo{
		value: value,
	}
	if c.cabeza == nil {
		c.cabeza = newNode
		c.cola = newNode
	} else {
		c.cola.next = newNode
		c.cola = newNode
	}
}

// Verificar datos

func (c *Cola) Size() int {
	actual := c.cabeza
	tamanio := 0
	for actual != nil {
		tamanio++
		actual = actual.next
	}
	return tamanio
}

func (c *Cola) Exist(carnet string) bool {
	nodoActual := c.cabeza
	for nodoActual != nil {
		if nodoActual.value.GetCarnet() == carnet {
			return true
		}
		nodoActual = nodoActual.next
	}
	return false
}

func (l *Cola) Print() {

	if l.cabeza == nil {
		fmt.Println("Cola vacía")
		return
	}
	nodoActual := l.cabeza

	for nodoActual != nil {
		nodoActual.value.Print()
		nodoActual = nodoActual.next
		fmt.Printf("%s\n", strings.Repeat("-", 45))
	}
}

// Eleminar datos

func (c *Cola) Dequeue() *est.Estudiante {
	if c.cabeza == nil {
		return nil
	}
	value := c.cabeza.value
	c.cabeza = c.cabeza.next
	if c.cabeza == nil {
		c.cola = nil
	}
	return value
}

func (c *Cola) Remove(index int) *est.Estudiante {
	if index < 0 || c.cabeza == nil {
		return nil
	}

	if index == 0 {
		return c.Dequeue()
	}

	node := c.cabeza
	for i := 0; i < index-1 && node != nil; i++ {
		node = node.next
	}

	if node == nil || node.next == nil {
		return nil
	}

	value := node.next.value
	node.next = node.next.next

	if node.next == nil {
		c.cola = node
	}

	return value
}

func (c *Cola) Get(index int) *est.Estudiante {
	if index < 0 {
		return nil
	}

	i := 0
	current := c.cabeza

	for current != nil {
		if i == index {
			return current.value
		}

		current = current.next
		i++
	}

	return nil
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
	fmt.Printf("Archivo %s creado exitosamente\n", nombre_archivo)
}

func ejecutar(nombre_imagen string, archivo_dot string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo_dot).Output()
	mode := 0777
	_ = ioutil.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
}

func (l *Cola) Graficar(nombreArchivo string) {

	nombre_archivo_dot := fmt.Sprintf("./%s.dot", nombreArchivo)
	nombre_imagen := fmt.Sprintf("%s.jpg", nombreArchivo)
	texto := "digraph cola{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = rectangle];\n"
	auxiliar := l.cabeza
	for i := 0; i < l.Size(); i++ {
		texto += fmt.Sprintf("nodo%s[label=\"%s\\n%s %s\"];\n", strconv.Itoa(i), auxiliar.value.GetCarnet(), auxiliar.value.GetNombre(), auxiliar.value.GetApellido())
		auxiliar = auxiliar.next
	}
	for i := 0; i < l.Size()-1; i++ {
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(i+1) + ";\n"
	}
	texto += "}"

	crearArchivoDot(nombre_archivo_dot)
	escribirArchivoDot(texto, nombre_archivo_dot)
	ejecutar(nombre_imagen, nombre_archivo_dot)
}
