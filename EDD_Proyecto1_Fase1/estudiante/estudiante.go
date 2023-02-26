package estudiante

import "fmt"

type Estudiante struct {
	nombre   string
	carnet   string
	password string
}

func New(nombre string, carnet string, password string) *Estudiante {
	e := new(Estudiante)
	e.nombre = nombre
	e.carnet = carnet
	e.password = password
	return e
}

func (e *Estudiante) Print() {
	fmt.Println("Nombre: ", e.nombre)
	fmt.Println("Carnet: ", e.carnet)
	fmt.Println("Contraseña: ", e.password)
}

func (e *Estudiante) GetNombre() string {
	return e.nombre
}
func (e *Estudiante) GetPassword() string {
	return e.password
}

func (e *Estudiante) GetCarnet() string {
	return e.carnet
}
