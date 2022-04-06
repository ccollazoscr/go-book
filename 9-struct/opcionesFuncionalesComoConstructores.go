package main

import "time"

type Estudiante struct {
	nombre     string
	nacimiento time.Time
	descuento  bool
}

type Opcion func(*Estudiante)

func Nombre(nombre string) Opcion {
	return func(stud *Estudiante) {
		stud.nombre = nombre
	}
}
func Nacimiento(nacimiento time.Time) Opcion {
	return func(stud *Estudiante) {
		stud.nacimiento = nacimiento
	}
}
func Descuento(descuento bool) Opcion {
	return func(stud *Estudiante) {
		stud.descuento = descuento
	}
}

func NewEstudiante(opciones ...Opcion) Estudiante {
	stud := Estudiante{
		nombre: "desconocido",
	}
	for _, opcion := range opciones {
		opcion(&stud)
	}
	return stud
}
