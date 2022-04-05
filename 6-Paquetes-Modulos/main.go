package main

import (
	"fmt"

	"github.com/mariomac/analizador"
	"macias.info/go/ejemplo/cubo"
	"macias.info/go/ejemplo/hola"
)

func main() {
	fmt.Print("Cómo te llamas?: ")
	var nombre string
	fmt.Scanln(&nombre)
	fmt.Println(hola.ConNombre(nombre))

	analizador.PrintEstadistica(nombre)

	fmt.Println("**********************")
	fmt.Println("El volumen de un cubo de lado 3.5 es", cubo.Volumen(3.5))
}
