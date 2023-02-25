package estudiante

import "fmt"

type Estudiante struct {
	nombre   string
	apellido string
	carnet   string
	password string
}

func New(nombre string, apellido string, carnet string, password string) *Estudiante {
	e := new(Estudiante)
	e.nombre = nombre
	e.apellido = apellido
	e.carnet = carnet
	e.password = password
	return e
}

func (e *Estudiante) Print() {
	fmt.Println("Nombre: ", e.nombre)
	fmt.Println("Apellido: ", e.apellido)
	fmt.Println("Carnet: ", e.carnet)
	fmt.Println("Contraseña: ", e.password)
}

func (e *Estudiante) GetNombre() string {
	return e.nombre
}
func (e *Estudiante) GetApellido() string {
	return e.apellido
}
func (e *Estudiante) GetPassword() string {
	return e.password
}

func (e *Estudiante) GetCarnet() string {
	return e.carnet
}
