package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan int)

	//Va a enviar numeros
	go senderNumber(chanel, 10)

	go receiver1(chanel)

	go receiver2(chanel)

	// Esperamos un poco para que las goroutines terminen su trabajo
	time.Sleep(35 * time.Second)
}

// Envía nros desde 1 a maxnro
func senderNumber(ch chan<- int, maxnro int) {
	for i := 1; i <= maxnro; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 200)
	}
	close(ch)
}

// Recibo el número y lo multiplico por "nrovar";
func receiver1(ch <-chan int) {
	nrovar := 10
	for num := range ch {
		result := num * nrovar
		fmt.Print("Receiver 2) ")
		fmt.Printf("Recibido: %d x %d = %d\n", num, nrovar, result)
		fmt.Println("------------------------------------------")
	}
}

// Recibo el número y lo sumo por "nrovar";
func receiver2(ch <-chan int) {
	nrovar := 10
	for num := range ch {
		result := num + nrovar
		fmt.Print("Receiver 2) ")
		fmt.Printf("Recibido: %d + %d = %d\n", num, nrovar, result)
		fmt.Println("------------------------------------------")
	}
}
