package main

import (
	"fmt"
	"time"
)

/*Multiples receptores*/
func main() {
	dulces := make(chan string, 10)
	go Engullidor("Marcos", dulces)
	go Engullidor("Aina", dulces)
	go Engullidor("Judit", dulces)

	dulces <- "Donut"
	time.Sleep(time.Second)
	dulces <- "Cruasan"
	time.Sleep(time.Second)
	dulces <- "Ensaimada"
	time.Sleep(time.Second)
	dulces <- "Ensalada"
	time.Sleep(time.Second)
}

/*dulces es un canal de solo lectura. Esta función se queda esperando a que le lleguen datos al canal para procesarlo. No se sabe a cual de
  los engullidores le será antendida*/
func Engullidor(nombre string, dulces <-chan string) {
	for dulce := range dulces {
		fmt.Println(nombre, "come", dulce)
	}
}
