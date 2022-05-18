package main

import (
	"fmt"
	"time"
)

/*Cancelando lecturas después de un tiempo de espera*/
func main() {
	//Para crear un canal se requier el uso de la función make
	ch := make(chan int)

	go func() {
		fmt.Println("Calculando respuesta a la pregunta de la vida")
		time.Sleep(15 * time.Second)
		ch <- 42
	}()

	fmt.Println("Esperando")
	select {
	case ret := <-ch:
		fmt.Println("Recibido", ret)
	case <-time.After(2 * time.Second):
		fmt.Println("Error: tiempo de espera agotado")
	}
}
