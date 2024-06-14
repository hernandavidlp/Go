package main

import (
	"fmt"
	"strings"
)

type Player struct {
	Name string
	Id   int
}

type Equipo struct {
	//Va a guardar el Id para poder seguir incrementando con nuevos jugadores
	IdGeneral int
	Jugadores []Player
}

func main() {
	miEquipo := Equipo{}
	miEquipo.IdGeneral = 1001
	//Carga de los jugadores por primera vez
	miEquipo.setearJugadores()
	//Muestra el equipo completo
	miEquipo.printPlayers()
	//Quita los 7 primeros jugadores
	miEquipo.deleteSevenPlayers()
	//Muestra el equipo con los 7 faltantes
	miEquipo.printPlayers()

	//Agrego los 7 nuevos jugadores
	miEquipo.addPlayer(limpiarCaracteres("Pedro Picapiedra"))
	miEquipo.addPlayer(limpiarCaracteres("Leon Gieco"))
	miEquipo.addPlayer(limpiarCaracteres("Walter Giardino"))
	miEquipo.addPlayer(limpiarCaracteres("Celine Dion"))
	miEquipo.addPlayer(limpiarCaracteres("Brais Moure"))
	miEquipo.addPlayer(limpiarCaracteres("Héctor De Leon"))
	miEquipo.addPlayer(limpiarCaracteres("Marcela Pagano"))

	//Muestra el equipo con los 7 aregados
	miEquipo.printPlayers()
}

// Permite agregar jugadores
func (e *Equipo) addPlayer(nombre string) {
	newPlayer := Player{Name: nombre, Id: e.IdGeneral}
	//Agrego el nuevo jugador al Slice
	e.Jugadores = append(e.Jugadores, newPlayer)
	//Incremento el IdGeneral para proximas agregaciones
	e.IdGeneral++
}

// Segun consigna se deben eliminar los 7 primeros jugadores
func (e *Equipo) deleteSevenPlayers() {
	e.Jugadores = e.Jugadores[7:]
}

// Imprime el equipo completo
func (e *Equipo) printPlayers() {
	fmt.Println("IMPRIME LA LISTA DE JUGADORES")
	fmt.Println("=============================")
	for i := 0; i < len(e.Jugadores); i++ {
		fmt.Printf("%d - %s\n", e.Jugadores[i].Id, e.Jugadores[i].Name)
	}
	fmt.Println("----- ----- ----- -----")
}

// Función que sirve para setear por primera vez todos los jugadores
func (e *Equipo) setearJugadores() {
	e.addPlayer("John Smith")
	e.addPlayer("David Johnson")
	e.addPlayer("Michael Garcia")
	e.addPlayer("Sarah Williams")
	e.addPlayer("Daniel Martinez")
	e.addPlayer("Emily Brown")
	e.addPlayer("James Rodriguez")
	e.addPlayer("Sophia Lee")
	e.addPlayer("Lucas Oliveira")
	e.addPlayer("Olivia Taylor")
	e.addPlayer("Mateo Hernandez")
	e.addPlayer("Emma Wilson")
	e.addPlayer("Gabriel Silva")
}

func limpiarCaracteres(texto string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
			return r
		}
		return -1

	}, texto)
}
