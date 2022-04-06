package main

import (
	"fmt"
	"time"
)

type Cuboide struct {
	Ancho    float64
	Alto     float64
	Profundo float64
}

/*Cuando NO se modifica los valores internnos NO es necesario pasarlos como referencia*/
func (c Cuboide) String() string {
	return fmt.Sprintf("%v * %v * %v", c.Ancho, c.Alto, c.Profundo)
}

/*Cuando NO se modifica los valores internnos NO es necesario pasarlos como referencia*/
func (c Cuboide) Volumen() float64 {
	return c.Ancho * c.Alto * c.Profundo
}

/*Cuando SI se modifica los valores internnos es necesario pasarlos como referencia*/
func (c *Cuboide) Redimensiona(an, al, pr float64) {
	c.Ancho = an
	c.Alto = al
	c.Profundo = pr
}

func main() {
	//c1 := Cuboide{} //forma de inicializar
	//var c2 Cuboide //Forma 2 de inicializar

	c := Cuboide{
		Ancho:    2,
		Alto:     3,
		Profundo: 2,
	}
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%#v\n", c)

	punterosAStruct()
	usoDeMetodosConStruct()

	testStructIncrustadas()
	estructuraVacia()
	opcionesFuncionalesComoConstructores()
}

func usoDeMetodosConStruct() {
	fmt.Println("********************")
	fmt.Println("usoDeMetodosConStruct")
	c := Cuboide{
		Ancho:    2,
		Alto:     3,
		Profundo: 2,
	}
	fmt.Printf("Cuboide %v. Volumen %v\n", c, c.Volumen())
	c.Redimensiona(1, 2, 3)
	fmt.Printf("Cuboide %v. Volumen %v\n", c, c.Volumen())
}

func punterosAStruct() {
	fmt.Println("********************")
	fmt.Println("punterosAStruct")
	v := Cuboide{}
	var p *Cuboide = &v
	p.Ancho = 1
	p.Alto = 2
	p.Profundo = 3

	fmt.Println(v)
	fmt.Println(&p)
}

func testStructIncrustadas() {
	fmt.Println("********************")
	fmt.Println("testStructIncrustadas")
	g := Aparato{}
	g.Nombre = "Reproductor MP3"
	g.Precio = 179
	g.AplicaDecuento(0.15)

	fmt.Printf("%+v\n", g)

	s := Telefono{}
	s.Nombre = "Nuevo Telefono"
	s.Precio = 600
	s.AplicaDecuento(0.15)

	fmt.Printf("%+v\n", s)

	//Instanciando un objeto
	s1 := Telefono{
		Aparato: Aparato{
			Nombre: "Test",
			Precio: 1000,
		},
		Pulgadas: 6,
		Bateria:  2400,
	}
	fmt.Printf("%+v\n", s1)
}

func estructuraVacia() {
	fmt.Println("********************")
	fmt.Println("estructuraVacia")
	//Instanciando una estructura vacia
	duplicados := []string{
		"Juan", "María", "Benito", "Juan", "Carlos",
		"Juan", "María", "Benito", "Juan", "Carlos",
		"Juan", "María", "Benito", "Juan", "Carlos",
	}

	unicos := map[string]struct{}{}
	for _, nombre := range duplicados {
		//Se asigna una instancia de estructura vacia
		unicos[nombre] = struct{}{}
	}

	fmt.Println("La lista de nombres únicos es:")
	for nombre := range unicos {
		fmt.Println(" -", nombre)
	}
}

func opcionesFuncionalesComoConstructores() {
	fmt.Println("********************")
	fmt.Println("opcionesFuncionalesComoConstructores")
	estud1 := NewEstudiante()
	estud2 := NewEstudiante(Nombre("Pedro"), Descuento(true))
	estud3 := NewEstudiante(Nombre("Juan"), Nacimiento(time.Date(2001, 10, 12, 0, 0, 0, 0, time.UTC)))

	fmt.Printf("%+v\n", estud1)
	fmt.Printf("%+v\n", estud2)
	fmt.Printf("%+v\n", estud3)
}
