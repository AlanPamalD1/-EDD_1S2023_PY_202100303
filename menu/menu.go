package main

import (
	"bufio"
	est "estudiante"
	"fmt"
	"io"
	lde "listadobleenlazada"
	"os"
	"os/exec"
	"strings"
)

var USER_ADMIN = "admin"
var PWD_ADMIN = "12345"

var COLA_PENDIENTES = lde.New()
var COLA_SISTEMA = lde.New()

func main() {
	//printMenuInicioSesion()
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
		fmt.Printf("*              %s             *\n", "5. Cerrar Sesión")
		fmt.Printf("%s\n", strings.Repeat("*", 45))
		fmt.Scanln(&option)
		switch option {
		case 1:
			//Ver estudiantes pendientes
			fmt.Printf("********** %s ***********\n", "Estudiantes Pendientes")
			COLA_PENDIENTES.Print()
			fmt.Printf("%s\n", strings.Repeat("*", 45))
		case 2:
			//Ver estudiantes del sistema
			fmt.Printf("********** %s ***********\n", "Estudiantes en Sistema")
			COLA_SISTEMA.Print()
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
			COLA_PENDIENTES.AddLast(objeto)
			fmt.Println("Estudiante agregado a lista de espera...")
		case 4:
			//carga masiva
		case 5:
			//Cerrar sesión
			fmt.Println("Cerrando sesión")
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
			fmt.Println("Saliendo del sistema ...")
			fmt.Println("Hasta la próxima")
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

func openImage() {
	path, _ := exec.LookPath("display")
	cmd, err := exec.Command(path, "lista.jpg").Output()
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("resultado: ", cmd)
}
