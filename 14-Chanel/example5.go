package main

import "fmt"

/*Sincronización mediante canales*/
func main() {
	espera := TareaAsincrona()
	<-espera

	fmt.Println("Programa finalizado")
}

func TareaAsincrona() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		fmt.Println("Haciendo alguna cosa en paralelo")
		for i := 0; i < 3; i++ {
			fmt.Println(i, "...")
		}
		fmt.Println("finaliza tarea en paralelo")

		close(ch)
	}()
	return ch
}
