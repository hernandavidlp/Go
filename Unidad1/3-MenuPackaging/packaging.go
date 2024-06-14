package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var countNewLaborer int
var countDeleteLaborer int
var Users []User
var UserLogin User
var Movimientos []string

type User struct {
	UserName string
	UserPass string
	UserRol  string
}

// Limpia la consola
func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Carga de datos el array Users con los 2 tipos de Usuarios
func setDataUsers() {
	Users = append(Users, User{UserName: "admin", UserPass: "root", UserRol: "Administrador"})
	Users = append(Users, User{UserName: "seeker223", UserPass: "seekr", UserRol: "Supervisor"})
}

func main() {
	setDataUsers()
	menu()
}

func login() {
	fmt.Println("====== LOGIN ======")
	fmt.Print("USUARIO: ")
	var username string
	fmt.Scanln(&username)
	fmt.Println()
	fmt.Print("PASSWORD: ")
	var userpass string
	fmt.Scanln(&userpass)
	for i := 0; i < len(Users); i++ {
		if Users[i].UserName == username && Users[i].UserPass == userpass {
			UserLogin = Users[i]
			break
		}
	}
	if !isEmpty(UserLogin) {
		if UserLogin.UserRol[0] == 'A' {
			submenuAdmin()
		} else {
			subMenuSupervisor()
		}
	} else {
		fmt.Println("No se encontró un usuario con ese nombre y contraseña...")
		login()
	}
}

func isEmpty(user User) bool {
	return user == User{}
}

func menu() {
	var exit bool = false
	for exit == false {
		fmt.Println()
		fmt.Println("=================================")
		fmt.Println("=  1 - Iniciar Sesión           =")
		fmt.Println("=  2 - Terminar programa        =")
		fmt.Println("=================================")
		var option string
		fmt.Scanln(&option)
		switch option {
		case "1":
			clearConsole()
			login()
		case "2":
			exit = true
		default:
			fmt.Println("Presiona 1 o 2")
			fmt.Println()
			menu()
		}
	}
}

func submenuAdmin() {
	fmt.Println()
	fmt.Println("==================================")
	fmt.Println("=         ADMINISTRADOR          =")
	fmt.Println("==================================")
	fmt.Println("=  1 - Crear un Laborer          =")
	fmt.Println("=  2 - Eliminar un Laborer       =")
	fmt.Println("=  3 - Generar Archivo de Texto  =")
	fmt.Println("=  4 - Cambiar a SUPERVISOR      =")
	fmt.Println("==================================")
	fmt.Print("Nro:")
	var nro int
	fmt.Scanln(&nro)
	switch nro {
	case 1:
		crearLaborer()
	case 2:
		eliminarLaborer()
	case 3:
		makeTextFile()
	case 4:
		clearConsole()
		println("Ingresa la credenciales de un Supervisor para poder acceder")
		login()
	default:
		println("Ingresa un número válido")
	}
}

func subMenuSupervisor() {
	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("=              SUPERVISOR               =")
	fmt.Println("=========================================")
	fmt.Println("=  1 - Crear registro de Administrador  =")
	fmt.Println("=  2 - Cambiar a ADMINISTRADOR          =")
	fmt.Println("=========================================")
	var nro int
	fmt.Scanln(&nro)
	switch nro {
	case 1:
		printLaborers()
	case 2:
		clearConsole()
		println("Ingresa la credenciales de un Administrador para poder acceder")
		login()
	default:
		println("Ingresa un número válido")
	}
}
func printLaborers() {
	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("        Registro de Administrador        ")
	fmt.Println("=========================================")
	for _, movimiento := range Movimientos {
		fmt.Println(movimiento)
	}
	fmt.Println("=========================================")
}

func crearLaborer() {
	countNewLaborer++
	str := "Laborer Creado - " + strconv.Itoa(countNewLaborer)
	Movimientos = append(Movimientos, str)
	clearConsole()
	submenuAdmin()
}
func eliminarLaborer() {
	countDeleteLaborer++
	str := "Laborer Eliminado - " + strconv.Itoa(countDeleteLaborer)
	Movimientos = append(Movimientos, str)
	clearConsole()
	submenuAdmin()
}

func makeTextFile() {
	clearConsole()
	// Obtener la ruta del perfil del usuario
	userProfile := os.Getenv("USERPROFILE")
	if userProfile == "" {
		fmt.Println("No se pudo obtener la ruta del perfil del usuario.")
		return
	}

	// Construir la ruta completa al escritorio
	desktopPath := filepath.Join(userProfile, "Desktop", "TextGo.txt")

	// Crear el archivo en el escritorio
	file, err := os.Create(desktopPath)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	// Leer el texto desde la consola
	fmt.Println("Ingresa el texto (presiona Ctrl+D en Unix/Linux/Mac o Ctrl+Z en Windows para finalizar):")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\x04') // \x04 es el código de fin de transmisión (Ctrl+D)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Error al leer el texto:", err)
		return
	}

	// Escribir el texto en el archivo
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Archivo 'TextGo.txt' creado y escrito con éxito en el escritorio.\n\n")
	submenuAdmin()
}
