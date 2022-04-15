package main

import "fmt"

//Defer se utiliza para realizar la ejecución de una función al final de su ciclo.
//Se utiliza por ejemplo para cerrar conexiones a BD
func CambiarRegistro() {
	fmt.Println("Hola")
	defer func() {
		fmt.Println("Ejecución aplazada")
	}()
	fmt.Println("Adios")
}
