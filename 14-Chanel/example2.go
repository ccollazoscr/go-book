package main

import "fmt"

const nums = 3

/*Ejemplo de manejo de canales de solo escritura y lectura*/
func main() {
	ch := make(chan int)
	go Emisor(ch)
	Receptor(ch)
}

/*Canales de solo escritura*/
func Emisor(ch chan<- int) {
	for i := 1; i <= nums; i++ {
		ch <- i
		fmt.Println(i, "enviando correctamente")
	}
}

/*Canales de solo lectura*/
func Receptor(ch <-chan int) {
	for i := 1; i <= nums; i++ {
		num := <-ch
		fmt.Println("recibido", num)
	}
}
