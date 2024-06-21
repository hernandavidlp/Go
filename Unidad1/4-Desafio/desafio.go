package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type User struct {
	UserName string
	UserPass string
	UserRol  Rol
}

type Rol string

const (
	Administrativo Rol = "Administrativo"
	Supervisor     Rol = "Supervisor"
	Cajero         Rol = "Cajero"
)

type Users struct {
	Users []User
}

type Venta struct {
	Fecha        time.Time
	Usuario      User
	DetalleVenta []VentaDetalle
}

type VentaDetalle struct {
	Producto Producto
	Cantidad int
}

type Ventas struct {
	Ventas []Venta
}

type Producto struct {
	Id     int
	Nombre string
	Precio float32
	Stock  int
}

type Productos struct {
	Id        int
	Productos []Producto
}

// Agrega Productos
func (p *Productos) addPrduct(nombre string, precio float32, stock int) {
	newProduct := Producto{Id: p.Id, Nombre: nombre, Precio: precio, Stock: stock}
	p.Productos = append(p.Productos, newProduct)
	//Incremento el id a nivel global
	p.Id++
}

// Limpia la consola
func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Agrega usuarios
func (u *Users) addUser(name string, password string, rol Rol) {
	newUser := User{UserName: name, UserPass: password, UserRol: rol}
	u.Users = append(u.Users, newUser)
}

// Imprime el listado de usuarios y sus roles
func (u *Users) printUsers() {
	fmt.Println("LISTADO DE USUARIOS Y ROLES")
	fmt.Println("===========================")
	for i, user := range u.Users {
		fmt.Printf("%s)", i)
		fmt.Printf("Nombre Usuario: %s\n", user.UserName)
		fmt.Printf("           Rol: %s\n", user.UserRol)
		fmt.Println()
	}
	fmt.Println("=========== FIN ===========")
}

func (u *Users) login() {
	fmt.Println("====== LOGIN ======")
	fmt.Print("USUARIO: ")
	var username string
	fmt.Scanln(&username)
	fmt.Println()
	fmt.Print("PASSWORD: ")
	var userpass string
	fmt.Scanln(&userpass)
	for i := 0; i < len(u.Users); i++ {
		if u.Users[i].UserName == username && u.Users[i].UserPass == userpass {
			if u.Users[i].UserRol == Administrativo {
				submenuAdmin()
			} else {
				if u.Users[i].UserRol == Supervisor {
					subMenuSupervisor()
				} else {
					subMenuLector()
				}
			}
		}
	}
	fmt.Println("No se encontró un usuario con ese nombre y contraseña...")
	u.login()
}

func subMenuCajero() {

}

func (u *Users) menu() {
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
			u.login()
		case "2":
			exit = true
		default:
			fmt.Println("Presiona 1 o 2")
			fmt.Println()
			u.menu()
		}
	}
}

func submenuAdmin() {
	fmt.Println()
	fmt.Println("==================================")
	fmt.Println("=         ADMINISTRATIVO         =")
	fmt.Println("==================================")
	fmt.Println("=  1 - Listar productos          =")
	fmt.Println("=  2 - Cambiar precio            =")
	fmt.Println("=  3 - Generar Archivo de Texto  =")
	fmt.Println("=  4 - Cambiar de Usuario        =")
	fmt.Println("=  5 - Menu anterior             =")
	fmt.Println("==================================")
	fmt.Print("Opcion:")
	var nro int
	fmt.Scanln(&nro)
	switch nro {
	case 1:

	case 2:

	case 3:
		makeTextFile()
	case 4:
		clearConsole()
		println("Ingresa la credenciales de un Supervisor para poder acceder")
		login()
	case 5:
		os.Exit(1)
	default:
		println("Ingresa un número válido")
	}
}

func subMenuSupervisor() {
	fmt.Println()
	fmt.Println("=========================================")
	fmt.Println("=              SUPERVISOR               =")
	fmt.Println("=========================================")
	fmt.Println("=  1 - Imprimir listado de productos    =")
	fmt.Println("=  2 - Cambiar de Usuario               =")
	fmt.Println("=  3 - Menu anterior                    =")
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

func main() {
	usuarios := Users{}
	usuarios.addUser("Pedro", "pedro123", Administrador)
	usuarios.addUser("Lucia", "lucia123", Supervisor)
	usuarios.addUser("Sofia", "sofia123", Lector)
	usuarios.addUser("Lucas", "lucas123", Lector)

	productos := Productos{}
	productos.Id = 1
	productos.addPrduct()

	menu()
}

func menuAgregarProducto() {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("=         Agregar un producto          =")
	fmt.Println("========================================")
	fmt.Print("Nombre:")
	fmt.Println("=========================================")
}
