package main

type PH float32

func (p PH) Categoria() string {
	switch {
	case p < 7:
		return "ácido"
	case p > 7:
		return "básico"
	default:
		return "neutro"
	}
}
