package main

import "fmt"

type ErrorTv struct {
	Causa error
}

func (e ErrorTv) Error() string {
	return fmt.Sprint("Problema con la TV:", e.Causa)
}

func (e ErrorTv) Unwrap() error {
	return e.Causa
}

type ErrorComponente struct {
	Nombre string
}

func (c ErrorComponente) Error() string {
	return "fallo de un compnente:" + c.Nombre
}
