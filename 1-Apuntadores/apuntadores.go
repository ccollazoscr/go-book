package main

import "fmt"

func main() {
	i := 10    //Asignando una variable cualquiera
	var p *int //Se declara un apuntador de tipo int
	p = &i     //A p se le asigna la dirección de memoria de i
	a := *p    //a 'a' se le asigna el valor de p. 'a' queda apuntando a otro espacio de memoria porque se le asigno el valor mas no la misma referencia
	*p = 21    //p cambia su valor. Como apunta al mismo espacio de memoria de i, i también cambia el valor a 21

	fmt.Printf("Imprimiendo el valor de a: %d \n", a)
	fmt.Printf("Imprimiendo el valor de p: %d \n", *p)
	fmt.Printf("Imprimiendo el valor de i: %d \n", i)
	fmt.Println(p)
	fmt.Println(&i)

	fmt.Println("****************************************")
	valoresVsReferencia()
}

func valoresVsReferencia() {
	i := 0
	j := 0
	p1 := &i
	p2 := &j

	if p1 == p2 { //Como ambos son apuntadores, se comparan si apuntan a la misma dirección
		fmt.Println("p1 y p2 apuntan a la misma dirección")
	} else {
		fmt.Println("p1 y p2 apuntan a direcciones distintas")
	}

	if *p1 == *p2 { //Aqui se comparan los valores de los apuntandores
		fmt.Println("p1 y p2 apuntan a valores iguales")
	} else {
		fmt.Println("p1 y p2 apuntan a valores distintos")
	}
}
