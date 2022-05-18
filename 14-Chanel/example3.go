package main

import "fmt"

/*Bloqueo en la escritura: Canales con o sin bufer*/
func main() {
	//El código comentado falla porque se esta haciendo uso de un canal sin bufer. La aplicación se queda bloqueada porque a recibido no le llega ningun valor
	// ch := make(chan string)
	// ch <- "hola"
	// recibido := <-ch
	// fmt.Println(recibido)

	CanalConBuffer()
	CanalConBufferYCierre()
	IterandoCanalesConFor()
}

func CanalConBuffer() {
	fmt.Printf("\nCanalConBuffer\n")
	//Declaración de un canal con bufer
	ch := make(chan string, 5)
	ch <- "hola"
	recibido := <-ch
	fmt.Println(recibido)
}

func CanalConBufferYCierre() {
	fmt.Printf("\nCanalConBufferYCierre\n")
	//Declaración de un canal con bufer
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	//El valor de 4 y 5 mostraran los valores por defecto. En este caso 0
	for i := 0; i < 5; i++ {
		fmt.Println("recibiendo", <-ch)
	}
}

func IterandoCanalesConFor() {
	fmt.Printf("\nIterandoCanalesConFor\n")
	//Declaración de un canal con bufer
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	//El valor de 4 y 5 mostraran los valores por defecto. En este caso 0
	for num := range ch {
		fmt.Println("recibiendo", num)
	}
	fmt.Println("Canal cerrado. Fin!")
}
