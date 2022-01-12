package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func main() {
	helloMessage := "Hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	fmt.Printf("helloMessage: %T", helloMessage)
	fmt.Printf("\ncursos: %T", cursos)
	fmt.Println("")

	message = "Hola #{nombre} interpolado"
	normalFunction(message)
}
