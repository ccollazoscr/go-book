package main

import "fmt"

func funcion2() {
	panic("Ha sucedido algo realmente malo")
	fmt.Println("texto no mostrado funcion2")
}

func funcion1() {
	//Recover se utiliza para recuperarse de un panic que se haya lanzado
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recuperandose del pánico: ", p)
		}
	}()
	funcion2()
	fmt.Println("texto no mostrado funcion1 ")
}
