package main

import "fmt"

func isPalindromo(text string) bool {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	return textReverse == text
}

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[2] = 2
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[:4])

	// Append

	slice = append(slice, 7)
	fmt.Println(slice)

	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	sliceRange := []string{"hola", "que", "hace"}

	for i := range sliceRange {
		fmt.Println(i)
	}

	if isPalindromo("amor a roma") {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No palindromo")
	}
}
