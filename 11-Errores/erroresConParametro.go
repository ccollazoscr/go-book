package main

import "fmt"

type Comida string

const (
	ComidaPerro  = Comida("Comida de perro")
	ComidaHumano = Comida("Comida de perro")
	ComidaGato   = Comida("Comida de perro")
	ComidaPajaro = Comida("Comida de perro")
)

type NoHambre struct{}

func (n NoHambre) Error() string {
	return "el perro no tiene hambre"
}

type NoApetecible struct {
	Ofrecido Comida
}

func (d NoApetecible) Error() string {
	return "a los perros no les gusta " + string(d.Ofrecido)
}

type Perro struct {
	TieneHambre bool
}

func (d *Perro) Alimenta(f Comida) error {
	if !d.TieneHambre {
		return NoHambre{}
	}
	if f != ComidaPerro && f != ComidaHumano {
		return NoApetecible{Ofrecido: f}
	}
	fmt.Println("comiendo", f, ": ñam ñam ñam")
	d.TieneHambre = false
	return nil
}
