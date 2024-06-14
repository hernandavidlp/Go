package main

import "fmt"

type User struct {
	Nombre   string
	Password string
	Roles    []string
}

type Users struct {
	Usuarios []User
}

func main() {
	usuarios := Users{}
	usuarios.addUsers("admin", "admin")
}

func (u *Users) addUsers(nombre string, password string) {
	newUser := User{Nombre: nombre,
		Password: password}
	u.Usuarios = append(u.Usuarios, newUser)
}

// Cuenta la cantidad de vocales
func contarVocales(texto string) int {
	result := 0
	vocales := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	for _, r := range texto {
		if vocales[r] {
			result++
		}
	}
	return result
}

func enviarMensaje(message string) {
	fmt.Println("Enviado: ", message)
}
