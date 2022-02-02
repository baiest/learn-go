package main

import "fmt"

func message(text string, channel chan string) {
	channel <- text
}

func main() {
	channel := make(chan string, 2)

	channel <- "Mensaje 1"
	channel <- "Mensaje 2"

	//len -> cuantos tienen en el momento
	// cap -> cuantos puede almacenar
	fmt.Println(len(channel), cap(channel))

	// Range y Close

	close(channel)

	// channel <- "Mensaje 3" //No se puede

	for message := range channel {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1: ", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2: ", m2)
		}
	}
}
