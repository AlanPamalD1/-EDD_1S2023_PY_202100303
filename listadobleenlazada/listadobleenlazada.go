package listadobleenlazada

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
	next   *Nodo
	before *Nodo
	value  *est.Estudiante
}

type ListaDoble struct {
	cabeza *Nodo
	cola   *Nodo
}

// Constructores

func New() *ListaDoble {
	l := new(ListaDoble)
	l.cabeza = nil
	l.cola = nil
	return l
}

// Insertar datos

func (l *ListaDoble) AddFirst(estudiante *est.Estudiante) {
	nodo := &Nodo{
		value: estudiante,
	}
	if l.cabeza == nil {
		l.cabeza = nodo
		l.cola = nodo
	} else {
		nodo.next = l.cabeza
		l.cabeza.before = nodo
		l.cabeza = nodo
	}
}

func (l *ListaDoble) AddLast(estudiante *est.Estudiante) {
	nodo := &Nodo{
		value: estudiante,
	}
	if l.cola == nil {
		l.cabeza = nodo
		l.cola = nodo
	} else {
		nodo.before = l.cola
		l.cola.next = nodo
		l.cola = nodo
	}
}

func (l *ListaDoble) AddInPos(posicion int, estudiante *est.Estudiante) bool {
	if posicion < 0 {
		return false
	}

	if posicion == 0 {
		l.AddFirst(estudiante)
		return true
	}

	nodo := &Nodo{value: estudiante}
	actual := l.cabeza

	for i := 0; i < posicion-1; i++ {
		if actual == nil {
			return false
		}
		actual = actual.next
	}
	if actual == nil {
		l.AddLast(estudiante)
		return true
	}
	nodo.next = actual.next
	nodo.before = actual
	actual.next = nodo
	if nodo.next != nil {
		nodo.next.before = nodo
	} else {
		l.cola = nodo
	}
	return true
}

// Verificar datos

func (l *ListaDoble) Size() int {
	actual := l.cabeza
	tamanio := 0
	for actual != nil {
		tamanio++
		actual = actual.next
	}
	return tamanio
}

func (l *ListaDoble) Exist(carnet string) bool {
	nodoActual := l.cabeza
	for nodoActual != nil {
		if nodoActual.value.GetCarnet() == carnet {
			return true
		}
		nodoActual = nodoActual.next
	}
	return false
}

func (l *ListaDoble) Print() {

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

func (l *ListaDoble) PopFirst() *est.Estudiante {
	if l.cabeza == nil {
		return nil
	}
	estudiante := l.cabeza.value
	if l.cabeza == l.cola {
		l.cabeza = nil
		l.cola = nil
	} else {
		l.cabeza = l.cabeza.next
		l.cabeza.before = nil
	}
	return estudiante
}

func (l *ListaDoble) PopLast() *est.Estudiante {
	if l.cola == nil {
		return nil
	}
	estudiante := l.cola.value
	if l.cabeza == l.cola {
		l.cabeza = nil
		l.cola = nil
	} else {
		l.cola = l.cola.before
		l.cola.next = nil
	}
	return estudiante
}

func (l *ListaDoble) PopInPos(posicion int) *est.Estudiante {
	if posicion < 0 || l.cabeza == nil {
		return nil
	}
	if posicion == 0 {
		return l.PopFirst()
	}
	actual := l.cabeza
	for i := 0; i < posicion-1; i++ {
		if actual.next == nil {
			return nil
		}
		actual = actual.next
	}
	if actual.next == nil {
		return nil
	}
	estudiante := actual.next.value
	if actual.next == l.cola {
		l.cola = actual
	} else {
		actual.next.next.before = actual
	}
	actual.next = actual.next.next
	return estudiante
}

// Ordenamiento

func (l *ListaDoble) SortByName() {
	if l.cabeza == nil {
		return
	}
	ordenado := false
	for !ordenado {
		ordenado = true
		nodoActual := l.cabeza
		for nodoActual.next != nil {
			if nodoActual.value.GetNombre() > nodoActual.next.value.GetNombre() {
				// Intercambiar los nodos
				nodoSiguiente := nodoActual.next
				nodoAnterior := nodoActual.before
				nodoSiguiente.before = nodoAnterior
				if nodoAnterior != nil {
					nodoAnterior.next = nodoSiguiente
				} else {
					l.cabeza = nodoSiguiente
				}
				nodoActual.next = nodoSiguiente.next
				if nodoSiguiente.next != nil {
					nodoSiguiente.next.before = nodoActual
				} else {
					l.cola = nodoActual
				}
				nodoSiguiente.next = nodoActual
				nodoActual.before = nodoSiguiente
				nodoActual = nodoSiguiente
				ordenado = false
			}
			nodoActual = nodoActual.next
		}
	}
}
func (l *ListaDoble) SortByCarnet() {
	if l.cabeza == nil {
		return
	}
	ordenado := false
	for !ordenado {
		ordenado = true
		nodoActual := l.cabeza
		for nodoActual.next != nil {
			if nodoActual.value.GetCarnet() > nodoActual.next.value.GetCarnet() {
				// Intercambiar los nodos
				nodoSiguiente := nodoActual.next
				nodoAnterior := nodoActual.before
				nodoSiguiente.before = nodoAnterior
				if nodoAnterior != nil {
					nodoAnterior.next = nodoSiguiente
				} else {
					l.cabeza = nodoSiguiente
				}
				nodoActual.next = nodoSiguiente.next
				if nodoSiguiente.next != nil {
					nodoSiguiente.next.before = nodoActual
				} else {
					l.cola = nodoActual
				}
				nodoSiguiente.next = nodoActual
				nodoActual.before = nodoSiguiente
				nodoActual = nodoSiguiente
				ordenado = false
			}
			nodoActual = nodoActual.next
		}
	}
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

func (l *ListaDoble) Graficar(nombreArchivo string) {
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
		texto += fmt.Sprintf("nodo%s[label=\"{|%s\\n%s %s|}\"];\n", strconv.Itoa(i), auxiliar.value.GetCarnet(), auxiliar.value.GetNombre(), auxiliar.value.GetApellido())
		auxiliar = auxiliar.next
	}
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Size()-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"

	crearArchivoDot(nombre_archivo_dot)
	escribirArchivoDot(texto, nombre_archivo_dot)
	ejecutar(nombre_imagen, nombre_archivo_dot)
}
