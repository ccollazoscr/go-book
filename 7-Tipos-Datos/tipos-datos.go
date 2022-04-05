package main

import (
	"fmt"
	"math/rand"
)

func main() {
	tiposDatosPorciones()
	tiposDatosMapas()
	tiposFuncionales()
	receptoresFuncion()
	receptoresFuncionApuntadores()
	tiposDeDatosPseudoEnumerados()
}

func tiposDeDatosPseudoEnumerados() {
	fmt.Println("************************************")
	fmt.Println("Tipos pseudoenumerados")
	var pb PaloBaraja = Espadas
	fmt.Println("La selección de la baraja es: ", pb.string())
}

func receptoresFuncionApuntadores() {
	fmt.Println("************************************")
	fmt.Println("Receptores de función con apuntadores")
	var c ContadorApuntadores
	c.Incrementa()
	c.Incrementa()
	c.Incrementa()
	fmt.Println("valor:", c)
	c.Reinicia(77)
	fmt.Println("tras reinicio", c)

}

func receptoresFuncion() {
	fmt.Println("************************************")
	fmt.Println("Receptores de función")
	phs := []PH{
		PH(7), PH(1.2), PH(9),
	}
	for _, ph := range phs {
		fmt.Printf("Un ph == %v es %v\n", ph, ph.Categoria())
	}
}

func tiposDatosPorciones() {
	fmt.Println("************************************")
	fmt.Println("Tipos de datos a partir de porciones")
	type Año int
	type Nacimiento []Año

	censo := Nacimiento{1979, 1983, 1965}
	censo = append(censo, 1990)
	censo = append(censo, 1955)

	suma := Año(0)
	for _, ano := range censo {
		suma += ano
	}
	media := suma / Año(len(censo))
	fmt.Println("Año de nacimiento medio:", media)
}

func tiposDatosMapas() {
	fmt.Println("************************************")
	fmt.Println("Tipos de datos a partir de mapas")
	type Año int
	type NombreNacimiento map[string]Año

	artistas := NombreNacimiento{
		"Vincent":  1853,
		"Elvis":    1935,
		"Salvador": 1904,
	}
	artistas["Rick"] = 1966
	for nombre, año := range artistas {
		fmt.Printf("%s nacio en %d\n", nombre, año)
	}
}

type Generador func() int

func tiposFuncionales() {
	fmt.Println("************************************")
	fmt.Println("Tipos funcionales")
	cnt := Contador()
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	cnt = Contador() //Genera un nuevo contador
	fmt.Println(cnt())

	cnt = Contador()
	rnd := Aleatorio(456)
	for i := 0; i < 5; i++ {
		fmt.Println(GenerarTodo(Cero, cnt, rnd)) //No se pasan parentesis porque lo que se pretende es pasar la función más no el resultado
	}
}

func Cero() int {
	return 0
}

func Contador() Generador {
	cuenta := 0
	return func() int {
		cuenta++
		return cuenta
	}
}

func Aleatorio(semilla int64) Generador {
	rnd := rand.NewSource(semilla)
	return func() int {
		return int(rnd.Int63())
	}
}

func GenerarTodo(gens ...Generador) []int {
	nums := make([]int, 0, len(gens))
	for _, fg := range gens {
		nums = append(nums, fg())
	}
	return nums
}
