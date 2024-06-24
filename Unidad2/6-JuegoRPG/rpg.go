package main

import "fmt"

type Jugador struct {
	Nombre  string
	Ataque  int
	Defensa int
	Energia int
	Vidas   int
}

type Enemigo struct {
	Nombre  string
	Ataque  int
	Defensa int
	Energia int
}

type Mundo struct {
	Nombre string
}

func crearJugador(nombre string) *Jugador {
	nuevoJugador := Jugador{
		Nombre:  nombre,
		Ataque:  40,
		Defensa: 30,
		Energia: 100,
	}
	return &nuevoJugador
}

func crearEnemigo(option string) *Enemigo {
	enemigo := Enemigo{}
	switch option {
	case "1":
		enemigo = Enemigo{Nombre: "Jaguar Negro",
			Ataque:  30,
			Defensa: 30,
			Energia: 100}
	case "2":
		enemigo = Enemigo{Nombre: "Tiburon Blanco",
			Ataque:  45,
			Defensa: 20,
			Energia: 100}
	case "3":
		enemigo = Enemigo{Nombre: "Serpiente Cobra",
			Ataque:  35,
			Defensa: 10,
			Energia: 100}
	}
	return &enemigo
}

func crearBatalla(jugador *Jugador, optionEscenario string) {

}

func main() {
	fmt.Println("Comienza el Juego...")
	fmt.Println("====================")
	fmt.Println("Ponle un nombre al Jugador: (Sin espacios)")
	var nombreJugador string
	fmt.Scanln(&nombreJugador)

	player := crearJugador(nombreJugador)

	fmt.Println()

	fmt.Println("Escoge un lugar:")
	fmt.Println("======================")
	fmt.Println("= 1. Selva           =")
	fmt.Println("= 2. Mar             =")
	fmt.Println("= 3. Desierto        =")
	fmt.Println("======================")
	var option string
	fmt.Scanln(&option)
	switch option {
	case "1":
		fmt.Println("Lucharás en la selva!")
	case "2":
		fmt.Println("Lucharás en el mar!")
	case "3":
		fmt.Println("Lucharás en el desierto!")
	}
	crearBatalla(player, option)
}
