package main

import (
	"errors"
	"fmt"
	"strings"
)

type Libro struct {
	Titulo string
	Autor  string
	Genero string
	Status BookStatus
}

type Biblioteca struct {
	Libros []Libro
}

type BookStatus string

const (
	Available BookStatus = "Disponible"
	Borrowed  BookStatus = "Prestado"
)

func main() {
	lista := Biblioteca{}
	lista.agregarLibro("La Inteligencia Emocional", "Daniel Goleman", "Psicología")
	lista.agregarLibro("Desordenados", "Jorge Daniel", "Humor")
	lista.agregarLibro("Atrapado Sin Salida", "Alam Brado", "Policial")
	lista.agregarLibro("La Bala que dobló la esquina", "Sifuentes Lourdes", "Policial")

	//Actualiza el estado del libro
	lista.updateEstadoLibro("Atrapado Sin Salida", Borrowed)

	busqueda := lista.buscaLibros("la")
	for _, libro := range busqueda {
		fmt.Println(libro.Titulo)
	}

}

// Permite agregar libros a la lista
func (l *Biblioteca) agregarLibro(titulo, autor, Genero string) {
	nuevoLibro := Libro{
		Titulo: titulo,
		Autor:  autor,
		Genero: Genero,
		Status: Available,
	}
	l.Libros = append(l.Libros, nuevoLibro)
}

// Actualiza el estado de un libro
func (b *Biblioteca) updateEstadoLibro(titulo string, status BookStatus) error {
	for i, libro := range b.Libros {
		if libro.Titulo == titulo {
			b.Libros[i].Status = status
			return nil
		}
	}
	return errors.New("No se encontró el libro")
}

// Permite borrar un libro por medio del título
func (b *Biblioteca) borrarLibro(titulo string) error {
	for i, libro := range b.Libros {
		if libro.Titulo == titulo {
			b.Libros = append(b.Libros[:i], b.Libros[i+1:]...)
			return nil
		}
	}
	return errors.New("No se encontró el libro")
}

func (b *Biblioteca) buscaLibros(texto string) []Libro {
	var encontrados []Libro
	for _, libro := range b.Libros {
		if strings.Contains(libro.Titulo, texto) ||
			strings.Contains(libro.Autor, texto) {

			encontrados = append(encontrados, libro)
		}
	}
	return encontrados
}
