package main

import (
	"fmt"
	"time"
)

func contar(canal chan int) {
	x := 0
	for {
		canal <- x
		x++
	}
}

func imprimir(canal chan int) {
	var valor int
	for {
		valor = <-canal
		fmt.Println(valor)
		time.Sleep(time.Second)
	}
}

/*Ejemplo de sincronización entre gorutines usando canales
 */
func main() {
	canal := make(chan int)
	go contar(canal)
	go imprimir(canal)
	var fin string
	fmt.Scanln(&fin)
}
