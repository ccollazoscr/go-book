package main

import "fmt"

func main() {
	//Para crear un canal se requier el uso de la función make
	ch := make(chan string)

	go func() {
		//Forma de asignar un valor a un canal
		ch <- "Hola"
	}()

	//Forma de recibir datos de un canal
	recibido := <-ch
	fmt.Println("He recibido", recibido)
}
