package main

import (
	"fmt"
	"math/rand"
)

func main() {
	capitales := map[string]string{
		"España":      "Madrid",
		"Francia":     "Paris",
		"Reino unido": "Londres",
		"Japón":       "Tokio",
	}
	for pais, capital := range capitales {
		fmt.Printf("La capital de %s es %s\n", pais, capital)
	}

	//Obtener la cantidad de elementos de un mapa
	numero := len(capitales)
	fmt.Println("contabilizadas", numero, "capitales")

	conjuntos()
}

func conjuntos() {
	fmt.Println("Los números ganadores de la lotería son:")
	numero := map[int]struct{}{}
	for len(numero) < 6 {
		n := rand.Intn(49) + 1
		if _, ok := numero[n]; !ok {
			numero[n] = struct{}{}
			fmt.Println("El", n, "...")
		}
	}
	fmt.Println("Felicidades a los ganadores")
}
