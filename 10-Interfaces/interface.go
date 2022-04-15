package main

import (
	"fmt"
	"time"
)

//******Ejemplo 1
type Saludador interface {
	Saluda() string
}

type Perro struct{}

func (p Perro) Saluda() string {
	return "Guau!"
}

type Gato struct{}

func (p Gato) Saluda() string {
	return "Miau!"
}

//********Ejemplo 2
type ConversorMilisegundos interface {
	Milliseconds() int64
}

func MuestraMS(c ConversorMilisegundos) {
	fmt.Println("esto son ", c.Milliseconds(), "milisegundos")
}

//********Ejemplo 3
type Val struct{}

func (a Val) String() string {
	return "Implementa String() con un receptor por valor"
}

type Apt struct{}

func (a *Apt) String() string {
	return "Implementa String() con un receptor por apuntador"
}

//********Ejemplo 4
type PerroBravo struct{}

func (p *PerroBravo) Ladra() string {
	return "Guau!"
}

type Rana struct{}

func (r *Rana) Croa() string {
	return "Croac!"
}

type Canguro struct{}

func (c *Canguro) Salta() string {
	return "Boing!"
}

func main() {
	fmt.Println("Ejemplo 1")
	var sld Saludador
	sld = Perro{}
	fmt.Println(sld.Saluda())
	sld = Gato{}
	fmt.Println(sld.Saluda())

	fmt.Println("Ejemplo 2")
	t := 23 * time.Second
	MuestraMS(t)

	fmt.Println("Ejemplo 3")
	//Interfaces implementadas por valor
	var str1 fmt.Stringer = Val{} //Se puede usuar por valor o por referencia
	var str2 fmt.Stringer = &Val{}
	fmt.Println(str1.String())
	fmt.Println(str2.String())

	//Interfaces implementadas por apuntador
	var str3 fmt.Stringer = &Apt{}
	//var str4 fmt.Stringer = Apt{} -> Genera error en tiempo de compilación
	fmt.Println(str3.String())

	v := Val{}
	p := &v
	fmt.Println(p.String())

	fmt.Println("Ejemplo 4")
	animales := []interface{}{123, Canguro{}, "Un helicoptero", PerroBravo{}, Rana{}} //El tipo interface se le puede asignar cualquier valor
	for _, a := range animales {
		fmt.Printf("%#v", a)
		if p, ok := a.(PerroBravo); ok { //Verificamos si es del tipo perro. En caso afirmativo entramos al if
			fmt.Print(": ", p.Ladra())
		}
		if r, ok := a.(Rana); ok {
			fmt.Print(": ", r.Croa())
		}
		if c, ok := a.(Canguro); ok {
			fmt.Print(": ", c.Salta())
		}
		fmt.Println()
	}
	fmt.Println()
	for _, a := range animales {

		switch x := a.(type) {
		case Rana:
			fmt.Println("Una rana", x.Croa())
		case Canguro:
			fmt.Println("Un canguro", x.Salta())
		case PerroBravo:
			fmt.Println("Un perro", x.Ladra())
		default:
			fmt.Println("No se que es esto", x)
		}
	}
}
