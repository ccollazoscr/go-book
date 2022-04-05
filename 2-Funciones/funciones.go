package main

import "fmt"

func main() {
	fmt.Println("Retorno de multiples valores")
	max, min := MaxMin(3, 2)
	fmt.Printf("%d > %d\n", max, min)

	fmt.Println("Retorno de multiples valores nombrados")
	max, min = MaxMinNamed(2, 6)
	fmt.Printf("%d > %d\n", max, min)

	fmt.Println("Paso por valor")
	a := 3
	fmt.Println("a = ", a)
	Incrementa(a)
	fmt.Println("a = ", a)

	fmt.Println("Paso por referencia")
	fmt.Println("a = ", a)
	IncrementaPorReferencia(&a)
	fmt.Println("a = ", a)

	fmt.Println("*******************")
	fmt.Println("Literales en función - lambdas - clousures")
	contador := 0
	generador := func() int {
		contador++
		return contador
	}
	fmt.Println("Generador contador:", generador(), generador(), generador())

	fmt.Println("Ejemplo 2 - literales en función")
	var operador func(int, int) int
	operador = suma
	fmt.Println("suma 3+4=", operador(3, 4))
	operador = multiplica
	fmt.Println("multiplica 3*4 =", operador(3, 4))

	fmt.Println("Ejemplo 3 literales en función")
	func() {
		fmt.Println("Esto se invocará de inmediato pero en otro contexto")
	}()

	fmt.Println("********* Suma de multiples valores *********")
	fmt.Println("1+2+3+4+5=", sumaMultiplesValores(1, 2, 3, 4, 5))
}

func sumaMultiplesValores(valores ...int) int {
	var suma int
	for _, valor := range valores {
		suma = suma + valor
	}
	return suma
}

func suma(a, b int) int {
	return a + b
}
func multiplica(a, b int) int {
	return a * b
}

func Incrementa(a int) {
	a++
}

func IncrementaPorReferencia(a *int) {
	*a++
}

func MaxMin(a, b int) (int, int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}

//En esta fx no es necesario especiicar el return con los valores
func MaxMinNamed(a, b int) (max int, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}
