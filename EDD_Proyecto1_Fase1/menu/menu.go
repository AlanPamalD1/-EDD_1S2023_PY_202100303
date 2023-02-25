package main

import (
	"bufio"
	cl "cola"
	est "estudiante"
	"fmt"
	"io"
	lde "listadobleenlazada"
	"os"
	"os/exec"
	plreg "pilaregistro"
	"strings"
)

var USER_ADMIN = "admin"
var PWD_ADMIN = "admin"

var COLA_PENDIENTES = cl.New()
var LISTA_SISTEMA = lde.New()
var PILA_REGISTRO = plreg.New()

//Estudiantes de prueba

func main() {
	//printMenuInicioSesion()

	pruebaEst := est.New("Alan Rodrigo", "Pamal De León", "202100303", "12345")
	pruebaEst2 := est.New("Diego Eduardo", "Cifuentes dardon", "202100304", "54321")
	pruebaEst3 := est.New("Luis Alejandro", "Perez Borja", "202100305", "00000")
	pruebaEst4 := est.New("Sergio Alejandro", "García nose", "202100302", "panqueue")
	COLA_PENDIENTES.Add(pruebaEst4)
	COLA_PENDIENTES.Add(pruebaEst)
	COLA_PENDIENTES.Add(pruebaEst2)
	COLA_PENDIENTES.Add(pruebaEst3)
	printMenuAdmin()
}

func printMenuAdmin() {
loopMenu:
	for {
		option := 0
		fmt.Println()
		fmt.Printf("*** %s ***\n", "Dashboard Administrador - EDD GoDrive")
		fmt.Printf("*       %s       *\n", "1. Ver Estudiantes Pendientes")
		fmt.Printf("*       %s      *\n", "2. Ver Estudiantes del Sistema")
		fmt.Printf("*       %s       *\n", "3. Registrar Nuevo Estudiante")
		fmt.Printf("*       %s      *\n", "4. Carga Masiva de Estudiantes")
		fmt.Printf("*                %s                *\n", "5. Reportes")
		fmt.Printf("*              %s             *\n", "6. Cerrar Sesión")
		fmt.Printf("%s\n", strings.Repeat("*", 45))
		fmt.Scanln(&option)
		switch option {
		case 1:
			//Ver estudiantes pendientes
			fmt.Printf("********** %s ***********\n", "Estudiantes en Sistema")
			COLA_PENDIENTES.Print()
			fmt.Printf("%s\n", strings.Repeat("*", 45))
			aceptarEstudiantes()
		case 2:
			//Ver estudiantes del sistema
			fmt.Printf("********** %s ***********\n", "Estudiantes en Sistema")
			LISTA_SISTEMA.Print()
			fmt.Printf("%s\n", strings.Repeat("*", 45))
		case 3:
			//Registrar nuevo estudiante
			s := bufio.NewScanner(os.Stdin)
			var nombre, apellido, carnet, password string = "", "", "", ""

			for {
				fmt.Print("Ingrese el nombre: ")
				var err error
				nombre, err = scanString(s)
				if err == nil {
					nombre := strings.TrimSpace(nombre)
					min := 1
					if len(nombre) >= min {
						break
					}
					err = fmt.Errorf(
						"El nombre debe de tener por lo menos %d caracteres",
						min,
					)
				}
				fmt.Println("Error en entrada", err)
			}

			for {
				fmt.Print("Ingrese el apellido: ")
				var err error
				apellido, err = scanString(s)
				if err == nil {
					apellido := strings.TrimSpace(apellido)
					min := 1
					if len(apellido) >= min {
						break
					}
					err = fmt.Errorf(
						"El apellido debe de tener por lo menos %d caracteres",
						min,
					)
				}
				fmt.Println("Error en entrada", err)
			}

			fmt.Print("Ingrese el carnet: ")
			fmt.Scanln(&carnet)

			for {
				fmt.Print("Ingrese la contraseña: ")
				var err error
				password, err = scanString(s)
				if err == nil {
					password := strings.TrimSpace(password)
					min := 4
					if len(password) >= min {
						break
					}
					err = fmt.Errorf(
						"El apellido debe de tener por lo menos %d caracteres",
						min,
					)
				}
				fmt.Println("Error en entrada", err)
			}

			//Creacion Estudiante
			objeto := est.New(nombre, apellido, carnet, password)

			//Agregado lista espera
			COLA_PENDIENTES.Add(objeto)

			fmt.Println("Estudiante agregado a lista de espera ...")
		case 4:
			//carga masiva
		case 5:
			//Menu de reportes
			printMenuReportes()
		case 6:
			//Cerrar sesión
			fmt.Println("~~~~ Cerrando sesión ~~~~")
			break loopMenu
		default:
			fmt.Println("Ingrese una opcíon válida")
		}
	}

}

func aceptarEstudiantes() {
	option := 0

	if COLA_PENDIENTES.Size() == 0 {
		fmt.Println("Cola vacía")
	}
loop:
	for COLA_PENDIENTES.Size() > 0 {
		fmt.Printf("\n************** Estudiantes pendientes: %d ****************\n", (COLA_PENDIENTES.Size()))
		estudianteActual := COLA_PENDIENTES.Get(0)
		fmt.Printf("Estudiante actual: %s %s\n", estudianteActual.GetNombre(), estudianteActual.GetApellido())
		fmt.Println("1. Aceptar al Estudiante")
		fmt.Println("2. Rechazar al Estudiante")
		fmt.Println("3. Volver al Menu")
		fmt.Printf("%s\n", strings.Repeat("-", 55))
		fmt.Print("Elige una opción: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			//Aceptar estudiante
			COLA_PENDIENTES.Dequeue()
			LISTA_SISTEMA.AddFirst(estudianteActual)
			LISTA_SISTEMA.SortByCarnet()
			PILA_REGISTRO.Push(estudianteActual, true)
			fmt.Printf("Se agregó a %s %s al sistema \n", estudianteActual.GetNombre(), estudianteActual.GetApellido())
		case 2:
			//Rechazar
			COLA_PENDIENTES.Dequeue()
			PILA_REGISTRO.Push(estudianteActual, false)
			fmt.Printf("Se rechazó a %s %s\n", estudianteActual.GetNombre(), estudianteActual.GetApellido())
		case 3:
			//Salir
			fmt.Println("~~~~ Regresando ~~~~")
			break loop
		default:
			fmt.Println("Opción no válida")
			fmt.Println("~~~~ Regresando ~~~~")
			break loop
		}
	}

	fmt.Printf("%s\n", strings.Repeat("*", 57))
}

func printMenuReportes() {
loopMenu:
	for {
		option := 0
		fmt.Println()
		fmt.Printf("****** %s *******\n", "Área de Reportes - EDD GoDrive")
		fmt.Printf("*         %s          *\n", "1. Estudiantes Aceptados")
		fmt.Printf("*         %s          *\n", "2. Estudiantes en Espera")
		fmt.Printf("*         %s          *\n", "3. Reporte Administrador")
		fmt.Printf("*                 %s                   *\n", "4. JSON")
		fmt.Printf("*                %s                *\n", "5. Regresar")
		fmt.Printf("%s\n", strings.Repeat("*", 45))
		fmt.Scanln(&option)
		switch option {
		case 1:
			//lista estudiantes aceptados
			if LISTA_SISTEMA.Size() > 0 {
				nombreArchivo := "Lista_Sistema"
				LISTA_SISTEMA.Graficar(nombreArchivo)
				//openImage(nombreArchivo)
			} else {
				fmt.Println("Lista vacía")
			}
		case 2:
			//cola estudiantes en espera
			if COLA_PENDIENTES.Size() > 0 {
				nombreArchivo := "Cola_Espera"
				COLA_PENDIENTES.Graficar(nombreArchivo)
				//openImage(nombreArchivo)
			} else {
				fmt.Println("Cola vacía")
			}
		case 3:
			//reporte Administrador
			if PILA_REGISTRO.Size() > 0 {
				nombreArchivo := "Pila_Registro"
				PILA_REGISTRO.Graficar(nombreArchivo)
				//openImage(nombreArchivo)
			} else {
				fmt.Println("Pila vacía")
			}
		case 4:
			//reporte en JSON

		case 5:
			//Regresar al menu administrador
			fmt.Println("~~~~ Regresando ~~~~")
			break loopMenu
		default:
			fmt.Println("Ingrese una opcíon válida")
		}
	}
}

func printMenuInicioSesion() {
loopMenu:
	for {
		option := 0
		fmt.Println()
		fmt.Printf("**************** %s ****************\n", "EDD GoDrive")
		fmt.Printf("*             %s             *\n", "1. Iniciar Sesión")
		fmt.Printf("*            %s           *\n", "2. Salir del Sistema")
		fmt.Printf("%s\n", strings.Repeat("*", 45))
		fmt.Scanln(&option)
		switch option {
		case 1:
			//Iniciar Sesión
			s := bufio.NewScanner(os.Stdin)
			var usuario, pwrd string = "", ""
			for {
				fmt.Print("Ingrese el usuario: ")
				var err error
				usuario, err = scanString(s)
				if err == nil {
					usuario := strings.TrimSpace(usuario)
					min := 1
					if len(usuario) >= min {
						break
					}
					err = fmt.Errorf(
						"El usuario debe de tener por lo menos %d caracteres",
						min,
					)
				}
				fmt.Println("Error en entrada", err)
			}
			for {
				fmt.Print("Ingrese la contraseña: ")
				var err error
				pwrd, err = scanString(s)
				if err == nil {
					pwrd := strings.TrimSpace(pwrd)
					min := 1
					if len(pwrd) >= min {
						break
					}
					err = fmt.Errorf(
						"La contraseña debe de tener por lo menos %d caracteres",
						min,
					)
				}
				fmt.Println("Error en entrada", err)
			}

			if usuario == USER_ADMIN && pwrd == PWD_ADMIN {
				fmt.Println("Ha ingresado el administrador")
				printMenuAdmin()
			}
		case 2:
			//Salir del sistema
			fmt.Println("~~~~ Saliendo del sistema ~~~~")
			fmt.Println(" Hasta la próxima ")
			break loopMenu
		default:
			fmt.Println("Ingrese una opcíon válida")
		}
	}
}

func scanString(s *bufio.Scanner) (string, error) {
	if s.Scan() {
		return s.Text(), nil
	}
	err := s.Err()
	if err == nil {
		err = io.EOF
	}
	return "", err
}

func openImage(nombreArchivo string) {

	fmt.Printf("Abriendo el archivo %s.jpg\n", nombreArchivo)
	path, _ := exec.LookPath("jpg")
	nombreImagen := fmt.Sprintf("%s.jpg", nombreArchivo)
	fmt.Println("ruta ", path, nombreImagen)
	cmd, err := exec.Command(path, nombreImagen).Output()
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("resultado: ", cmd)
}
