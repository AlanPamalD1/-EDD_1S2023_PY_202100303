package main

import (
	"bufio"
	cl "cola"
	"encoding/csv"
	est "estudiante"
	"fmt"
	"io"
	lde "listadobleenlazada"
	"os"
	"path/filepath"
	plbtc "pilabitacora"
	plreg "pilaregistro"
	"strings"
	"syscall"
)

var USER_ADMIN = "admin"
var PWD_ADMIN = "admin"

var COLA_PENDIENTES = cl.New()
var LISTA_SISTEMA = lde.New()
var PILA_REGISTRO = plreg.New()

//Estudiantes de prueba

func main() {
	menuInicioSesion()
}

func menuAdmin() {
loopMenu:
	for {
		option := 0
		fmt.Println()
		fmt.Printf("*** %s ***\n", "Dashboard Administrador - EDD GoDrive")
		fmt.Printf("*       %s       *\n", "1. Ver Estudiantes pendientes")
		fmt.Printf("*       %s      *\n", "2. Ver Estudiantes del sistema")
		fmt.Printf("*       %s       *\n", "3. Registrar nuevo estudiante")
		fmt.Printf("*       %s      *\n", "4. Carga masiva de estudiantes")
		fmt.Printf("*                %s                *\n", "5. Reportes")
		fmt.Printf("*              %s             *\n", "6. Cerrar sesión")
		fmt.Printf("%s\n", strings.Repeat("*", 45))
		fmt.Scanln(&option)
		switch option {
		case 1:
			//Ver estudiantes pendientes
			if COLA_PENDIENTES.Size() > 0 {
				fmt.Printf("********** %s ***********\n", "Estudiantes en Sistema")
				aceptarEstudiantes()
				fmt.Printf("%s\n", strings.Repeat("*", 45))
			} else {
				fmt.Println("No hay estudiante pendientes")
			}
		case 2:
			if LISTA_SISTEMA.Size() > 0 {
				//Ver estudiantes del sistema
				fmt.Printf("********** %s ***********\n", "Estudiantes en Sistema")
				LISTA_SISTEMA.Print()
				fmt.Printf("%s\n", strings.Repeat("*", 45))
			} else {
				fmt.Println("No hay estudiante en el sistema")
			}
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

			nombre += " " + apellido

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
						"La contraseña debe de tener por lo menos %d caracteres",
						min,
					)
				}
				fmt.Println("Error en entrada", err)
			}

			//Creacion Estudiante
			objeto := est.New(nombre, carnet, password)

			//Agregado lista espera
			COLA_PENDIENTES.Add(objeto)

			fmt.Println("Estudiante agregado a lista de espera ...")
		case 4:
			//carga masiva
			ruta := ""
			fmt.Print("Ingrese la ruta del archivo: ")
			fmt.Scanln(&ruta)
			CargaMasivaEstudiantes(ruta)
		case 5:
			//Menu de reportes
			menuReportes()
		case 6:
			//Cerrar sesión
			fmt.Println("Cerrando sesión ...")
			break loopMenu
		default:
			fmt.Println("Ingrese una opcíon válida")
		}
	}
}

func aceptarEstudiantes() {
loop:
	for COLA_PENDIENTES.Size() > 0 {
		option := 0
		fmt.Printf("\n************** Estudiantes pendientes: %d ****************\n", (COLA_PENDIENTES.Size()))
		estudianteActual := COLA_PENDIENTES.Get(0)
		fmt.Printf("Estudiante actual: %s\n", estudianteActual.GetNombre())
		fmt.Println("1. Aceptar al estudiante")
		fmt.Println("2. Rechazar al estudiante")
		fmt.Println("3. Volver al menu")
		fmt.Printf("%s\n", strings.Repeat("-", 55))
		fmt.Print("Elige una opción: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			//Aceptar estudiante
			//Quitar de cola de pendientes
			COLA_PENDIENTES.Dequeue()
			//Agregar a lista de sistema
			pila_bitacora := plbtc.New()
			LISTA_SISTEMA.AddFirst(estudianteActual, pila_bitacora)
			//Ordenar datos
			LISTA_SISTEMA.SortByCarnet()
			//Agregar al registro de la administrador
			PILA_REGISTRO.Push(estudianteActual, true)
			fmt.Printf("Se agregó a %s al sistema \n", estudianteActual.GetNombre())
		case 2:
			//Rechazar
			COLA_PENDIENTES.Dequeue()
			PILA_REGISTRO.Push(estudianteActual, false)
			fmt.Printf("Se rechazó a %s\n", estudianteActual.GetNombre())
		case 3:
			//Salir
			fmt.Println("Regresando ...")
			break loop
		default:
			fmt.Println("Opción no válida")
			fmt.Println("Regresando ...")
			break loop
		}
	}

	fmt.Printf("%s\n", strings.Repeat("*", 57))
}

func menuReportes() {
loopMenu:
	for {
		option := 0
		fmt.Println()
		fmt.Printf("****** %s *******\n", "Área de Reportes - EDD GoDrive")
		fmt.Printf("*         %s          *\n", "1. Estudiantes en espera")
		fmt.Printf("*        %s         *\n", "2. Estudiantes registrados")
		fmt.Printf("*         %s          *\n", "3. Reporte administrador")
		fmt.Printf("*                 %s                   *\n", "4. JSON")
		fmt.Printf("*                %s                *\n", "5. Regresar")
		fmt.Printf("%s\n", strings.Repeat("*", 45))
		fmt.Scanln(&option)
		switch option {
		case 1:
			//lista estudiantes aceptados
			if LISTA_SISTEMA.Size() == 0 {
				fmt.Println("Lista vacía")
			}
			LISTA_SISTEMA.Graficar("Estudiantes registrados")
			//openImage("Estudiantes registrados")

		case 2:
			//cola estudiantes en espera
			if COLA_PENDIENTES.Size() == 0 {
				fmt.Println("Cola vacía")
			}
			COLA_PENDIENTES.Graficar("Estudiantes pendientes")
			//openImage("Estudiantes pendientes")

		case 3:
			//reporte Administrador
			if COLA_PENDIENTES.Size() == 0 {
				fmt.Println("Pila vacía")
			}
			PILA_REGISTRO.Graficar("Reporte administrador")
			//openImage("Reporte administrador")
		case 4:
			//reporte en JSON
			GenerarJsonSistema("Reporte")
		case 5:
			//Regresar al menu administrador
			fmt.Println("Regresando ...")
			break loopMenu
		default:
			fmt.Println("Ingrese una opcíon válida")
		}
	}
}

func menuInicioSesion() {
loopMenu:
	for {
		option := 0
		fmt.Println()
		fmt.Printf("**************** %s ****************\n", "EDD GoDrive")
		fmt.Printf("*             %s             *\n", "1. Iniciar sesión")
		fmt.Printf("*            %s           *\n", "2. Salir del sistema")
		fmt.Printf("%s\n", strings.Repeat("*", 45))
		fmt.Scanln(&option)
		switch option {
		case 1:
			//Iniciar Sesión
			s := bufio.NewScanner(os.Stdin)
			var usuario, pwrd string = "", ""

			fmt.Print("Ingrese el usuario: ")
			fmt.Scanln(&usuario)

			for {
				fmt.Print("Ingrese la contraseña: ")
				var err error
				pwrd, err = scanString(s)
				if err == nil {
					pwrd := strings.TrimSpace(pwrd)
					min := 4
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
			if usuario == USER_ADMIN {
				if pwrd == PWD_ADMIN {
					fmt.Println("Ha ingresado el administrador")
					menuAdmin()
				} else {
					fmt.Println("Contraseña incorrecta")
				}
			} else {
				//Busqueda usuario en lista doble enlazada usando carnet (usuario)
				busqueda_usuario := LISTA_SISTEMA.GetStudentByCarnet(usuario)
				if busqueda_usuario == nil {
					fmt.Println("Usuario no registrado con este carnet")
				} else {
					if busqueda_usuario.GetPassword() == pwrd {
						//Ingreso de usuario
						busqueda_bitacora := LISTA_SISTEMA.GetStackByCarnet(usuario)
						busqueda_bitacora.Push()
						menuUsuario(busqueda_usuario)
					} else {
						//Contraseña incorrecta
						fmt.Println("Contraseña incorrecta")
					}
				}
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

func menuUsuario(estudiante *est.Estudiante) {
loopMenu:
	for {
		option := 0
		fmt.Println()
		fmt.Printf("**************** %s ****************\n", "EDD GoDrive")
		fmt.Printf("*        %s         *\n", "1. Historial inicio sesión")
		fmt.Printf("*          %s          *\n", "2. Graficar mi historial")
		fmt.Printf("*              %s             *\n", "3. Cerrar Sesión")
		fmt.Printf("%s\n", strings.Repeat("*", 45))
		fmt.Scanln(&option)
		switch option {
		case 1:
			//Mostrar bitácora
			fmt.Printf("************ %s *************\n", "Historial Ingreso")
			busqueda_bitacora := LISTA_SISTEMA.GetStackByCarnet(estudiante.GetCarnet())
			busqueda_bitacora.Print()
			fmt.Printf("%s\n", strings.Repeat("*", 45))
		case 2:
			busqueda_bitacora := LISTA_SISTEMA.GetStackByCarnet(estudiante.GetCarnet())

			if busqueda_bitacora.Size() > 0 {
				busqueda_bitacora.Graficar("Bitacora")
				////openImage("Bitacora")
			} else {
				fmt.Println("No hay registro")
			}
		case 3:
			//Salir
			fmt.Println("Cerrando sesión ...")
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

	execFPath, _ := os.Executable()
	execPath := filepath.Dir(execFPath)

	var sI syscall.StartupInfo
	var pI syscall.ProcessInformation
	argv, _ := syscall.UTF16PtrFromString(os.Getenv("windir") + "\\system32\\cmd.exe /C " + fmt.Sprintf("\"%s\\%s.jpg\"", execPath, nombreArchivo))
	syscall.CreateProcess(nil, argv, nil, nil, true, 0, nil, nil, &sI, &pI)
}

func CargaMasivaEstudiantes(ruta string) {
	//abrir archivo
	f, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No se pudo abrir el archivo")
		return
	}
	//cerrar el archivo
	defer f.Close()
	//lectura csv
	csvReader := csv.NewReader(f)
	contador := 0
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No se pudo leer el CSV")
		}
		//lectura de lineas
		//saltar encabezado
		if contador > 0 {
			carnet := rec[0]
			nombre := rec[1]
			password := rec[2]
			estudiante_nuevo := est.New(nombre, carnet, password)
			COLA_PENDIENTES.Add(estudiante_nuevo)
		}
		contador += 1
	}
	fmt.Printf("Se agregó %d estudiantes a pendientes\n", (contador - 1))
}

func GenerarJsonSistema(nombre string) {
	LISTA_SISTEMA.GenerarArchivoJSON(nombre)
}
