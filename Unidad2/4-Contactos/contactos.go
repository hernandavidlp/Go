package main

import (
	"errors"
	"fmt"
	"strings"
)

type Contacto struct {
	Nombre    string
	Telefono  string
	Email     string
	Direccion string
}

type ListaContactos struct {
	Contactos []Contacto
}

func main() {
	lista := ListaContactos{}
	lista.addConctacto("Daniel Zarate", "+54376432114", "daniel@yahoo.com", "Los Pinos 1214")
	lista.addConctacto("Jaziel Torres", "+15676432480", "jaziel@outlook.com", "Chacabuco y Cerrito")
	lista.addConctacto("Rossana Diaz", "+24375137542", "rossana@protonmail.com", "Timbó 34")
	lista.addConctacto("Paola Lenur", "+57376478425", "paola@gmail.com", "El manantial 547")

	//Actualiza los datos de un contacto buscando por nombre
	lista.updateContactoPorNombre("Rossana Diaz", "+243751375487", "rossanadiaz@outlook.com", "Los Pinos 1214")

	//Actualiza los datos de un contacto buscando por telefono
	lista.updateContactoPorTelefono("Paola Lenurte", "+57376478425", "paola@gmail.com", "Catedral Centro 225")

	//Para verificar imprimo en pantalla
	fmt.Println("LISTADO DE CONTACTOS")
	fmt.Println("====================")
	for _, contacto := range lista.Contactos {
		fmt.Printf("   Nombre: %s\n", contacto.Nombre)
		fmt.Printf(" Telefono: %s\n", contacto.Telefono)
		fmt.Printf("    Email: %s\n", contacto.Email)
		fmt.Printf("Direccion: %s\n", contacto.Direccion)
		println("----- ------ ------")
	}

}

// Permite agregar contactos a la lista
func (lc *ListaContactos) addConctacto(nombre, telefono, email, direccion string) {
	nuevoContacto := Contacto{
		Nombre:    nombre,
		Telefono:  telefono,
		Email:     email,
		Direccion: direccion,
	}
	lc.Contactos = append(lc.Contactos, nuevoContacto)
}

// Permite la búsqueda de un Contacto por nombre o Nro de teléfono
func (lc *ListaContactos) buscaContactos(texto string) []Contacto {
	var encontrados []Contacto
	for _, contacto := range lc.Contactos {
		if strings.Contains(contacto.Nombre, texto) ||
			strings.Contains(contacto.Telefono, texto) {

			encontrados = append(encontrados, contacto)
		}
	}
	return encontrados
}

// Actualiza los datos de un contacto buscando por nombre
func (lc *ListaContactos) updateContactoPorNombre(nombre, telefono, email, direccion string) error {
	for i, contacto := range lc.Contactos {
		if contacto.Nombre == nombre {
			nuevoContacto := Contacto{Nombre: nombre, Telefono: telefono, Email: email, Direccion: direccion}
			lc.Contactos[i] = nuevoContacto
			return nil
		}
	}
	return errors.New("No se encontró ningún contacto")
}

// Actualiza los datos de un contacto buscando por nro de teléfono
func (lc *ListaContactos) updateContactoPorTelefono(nombre, telefono, email, direccion string) error {
	for i, contacto := range lc.Contactos {
		if contacto.Telefono == telefono {
			nuevoContacto := Contacto{Nombre: nombre, Telefono: telefono, Email: email, Direccion: direccion}
			lc.Contactos[i] = nuevoContacto
			return nil
		}
	}
	return errors.New("No se encontró ningún contacto")
}

// Permite borrar un contacto por medio del título
func (lc *ListaContactos) borrarContacto(nombre string) error {
	for i, contacto := range lc.Contactos {
		if contacto.Nombre == nombre {
			lc.Contactos = append(lc.Contactos[:i], lc.Contactos[i+1:]...)
			return nil
		}
	}
	return errors.New("No se encontró un contacto para borrar")
}
