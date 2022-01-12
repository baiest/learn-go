package main

import "fmt"

func areaCuadrado() {
	const baseCuadrado int = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	result := x + y
	fmt.Println("Suma", result)

	result = y - x
	fmt.Println("Resta", result)

	result = x * y
	fmt.Println("Multiplicación", result)

	result = y / x
	fmt.Println("División", result)

	result = y % x
	fmt.Println("Modulo", result)

	x++
	fmt.Println("Incremental", x)

	y--
	fmt.Println("Decremental", y)

}

func main() {
	//Declaracion de constantes
	const PI float64 = 3.14
	const PI2 = 3.1416
	fmt.Println("Hello world Go")
	fmt.Println("PI:", PI)
	fmt.Println("PI2:", PI2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero value
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
	areaCuadrado()
}
