package main

import "fmt"

func main() {
	fmt.Println("Ingresa un numero positivo para empezar:")
	var nro int32
	_, err := fmt.Scanln(&nro)
	if err != nil {
		fmt.Println(err)
	} else {
		fizzBuzz(nro)
	}
}

func fizzBuzz(nro int32) {
	for i := 1; i <= int(nro); i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}
		if i < int(nro) {
			fmt.Print(", ")
		}
	}
}
