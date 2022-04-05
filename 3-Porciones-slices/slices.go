package main

import "fmt"

func main() {
	//Declarando un slice. No se especifica el tamaño
	s := []int{1, 2, 3, 4} //slice
	v1 := [3]int{1, 2, 3}  //Vector
	v2 := [3]int{1, 2, 3}  //Vector

	fmt.Println(s)
	fmt.Println(v1)
	fmt.Println(v2)

	//Adicionando elementos a un slice
	var p []int
	p = append(p, 1, 2, 3)
	p = append(p, 4)
	fmt.Println(p)

	//Longitud vs capacidad
	var sl []int
	fmt.Printf("Longitud %v. Capacidad %v\n", len(sl), cap(sl))
	sl = append(sl, 1, 2, 3, 4)
	fmt.Printf("Longitud %v. Capacidad %v\n", len(sl), cap(sl))
	sl = append(sl, 5)
	fmt.Printf("Longitud %v. Capacidad %v\n", len(sl), cap(sl))

	//Es recomendable utilizar make si se conoce el tamaño del slice
	altaCapacidad := make([]float32, 0, 2048)
	fmt.Printf("Longitud %v. Capacidad %v\n", len(altaCapacidad), cap(altaCapacidad))

	//Copia de slices con copy
	original := []int{1, 2, 3, 4, 5}
	copia := make([]int, len(original))

	n := copy(copia, original)
	fmt.Println(n, "números copiados", copia)

	//Vectores y slices en parámetros de función. Los vectores trabajan por valor, los slices trabajan por referencia
	fmt.Println("Vectores y slices en parámetros de función")
	v := [3]int{1, 2, 3}
	p2 := []int{1, 2, 3}
	cero(v, p2)
	fmt.Println("vector: ", v)
	fmt.Println("slice: ", p2)

	//Recorrer vectores y porciones
	fmt.Println("Recorrer vectores y porciones")
	ciudades := []string{"Tokio", "Paris", "Roma", "El cairo"}
	for indice := 0; indice < len(ciudades); indice++ {
		fmt.Printf("[%d] %s\n", indice, ciudades[indice])
	}
	//La forma más optima de recorrerlos
	fmt.Println("Recorriendolos con range")
	for indice, ciudad := range ciudades {
		fmt.Printf("[%d] %s\n", indice, ciudad)
	}

	//Creando vistas desde los slices. El indice inicial es incluyente, el final es excluyente
	fmt.Println("Creando vistas desde los slices")
	base := []int{1, 0, 3, 4, 5}

	vista1 := base[1:3] //incluye elementos de la posición 1 y 2
	fmt.Println("vista 1: ", vista1)
	vista1[0] = 2
	fmt.Println("base: ", base)

	vista2 := base[2:] //Desde el indice 2 hasta el final
	fmt.Println("vista 2: ", vista2)

	vista3 := base[:3] //Desde el indice 0 hasta el indice 2
	fmt.Println("vista 3: ", vista3)

	vista4 := base[:] //Vista de inicio a fin
	fmt.Println("vista 4: ", vista4)

	//Haciendo uso del operador difusor. Este operador se utiliza para pasar un slice como parámetro a una función que recibe multiples valores ...
	fmt.Println("Usando el operador difusor")
	valoresSumar := []int{1, 2, 3, 4, 5}
	suma := sumaMultiplesValores(valoresSumar...) //Usando el operador difusor
	fmt.Println("sumando", valoresSumar, "da igual a", suma)
}

func sumaMultiplesValores(valores ...int) int {
	var suma int
	for _, valor := range valores {
		suma = suma + valor
	}
	return suma
}

func cero(vec [3]int, porc []int) {
	vec[0] = 0
	if len(porc) > 0 {
		porc[0] = 0
	}
}
