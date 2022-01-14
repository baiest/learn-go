package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// And

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	//Or
	if valor1 == 3 || valor2 == 2 {
		fmt.Println("Es verdad")
	}

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No hay condicion")
	}
}
