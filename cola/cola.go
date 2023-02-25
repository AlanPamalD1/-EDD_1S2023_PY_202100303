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
	next   *Nodo
	before *Nodo
	value  *est.Estudiante
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

func (l *Cola) Add(estudiante *est.Estudiante) {
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

// Verificar datos

func (l *Cola) Size() int {
	actual := l.cabeza
	tamanio := 0
	for actual != nil {
		tamanio++
		actual = actual.next
	}
	return tamanio
}

func (l *Cola) Exist(carnet string) bool {
	nodoActual := l.cabeza
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

func (l *Cola) Pop() *est.Estudiante {
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

// Ordenamiento

func (l *Cola) SortByName() {
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
func (l *Cola) SortByCarnet() {
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

func (l *Cola) Graficar(nombreArchivo string) {
	fmt.Println("Impresion")
	nombre_archivo_dot := fmt.Sprintf("./%s.dot", nombreArchivo)
	nombre_imagen := fmt.Sprintf("%s.jpg", nombreArchivo)
	texto := "digraph cola{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	auxiliar := l.cabeza
	contador := 0
	for i := 0; i < l.Size(); i++ {
		//texto = texto + fmt.Sprintf("nodo%s[label=\"{|%s %s|}\"];\n", strconv.Itoa(i), auxiliar.value.GetCarnet(), auxiliar.value.GetNombre())
		texto += fmt.Sprintf("nodo%s[label=\"{|%s\\n%s|}\"];\n", strconv.Itoa(i), auxiliar.value.GetCarnet(), auxiliar.value.GetNombre())
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
