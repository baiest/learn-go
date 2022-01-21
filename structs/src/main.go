package main

import (
	"fmt"
)

type car struct {
	marca string
	Color string
}

func PrintMessage() {
	fmt.Println("Hola")
}

func main() {
	var myCar car
	myCar.marca = "Ferrari"
	fmt.Println(myCar)
}
