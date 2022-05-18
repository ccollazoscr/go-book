package main

import (
	"context"
	"fmt"
	"time"
)

/*Cancelando mediante contextos*/
func main() {
	ctx, cancela := context.WithCancel(context.Background())

	fmt.Println("Un mensaje se mostrará en 5 segundos")
	go func() {
		fmt.Println("Pulsa intro para cancelar el mensaje")
		fmt.Scanf("\n")
		cancela()
	}()

	MuestraRetardada(ctx, "Hola!")
	fmt.Println("Finalizando!")
}

func MuestraRetardada(ctx context.Context, msg string) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println(msg)
	case <-ctx.Done():
		//Proceso interrumpido. La fx continua
	}
}
