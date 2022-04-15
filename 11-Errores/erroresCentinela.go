package main

import "errors"

var (
	ErrYaExiste      = errors.New("clave ya existe")
	ErrNoCapacidad   = errors.New("no hay capacidad")
	ErrClaveInvalida = errors.New("clave invalida")
)

type Almacen struct {
	capacidad int
	elementos map[string]interface{}
}

func NewAlmacen(capacidad int) Almacen {
	return Almacen{
		capacidad: capacidad,
		elementos: map[string]interface{}{},
	}
}

func (s *Almacen) Guarda(clave string, valor interface{}) error {
	if clave == "" {
		return ErrClaveInvalida
	}
	if len(s.elementos) >= s.capacidad {
		return ErrNoCapacidad
	}
	if _, ok := s.elementos[clave]; ok {
		return ErrYaExiste
	}
	s.elementos[clave] = valor
	return nil
}
