package main

type PaloBaraja int

const (
	Espadas PaloBaraja = 0
	Bastos  PaloBaraja = 1
	Oros    PaloBaraja = 2
	Copas   PaloBaraja = 3
)

func (f PaloBaraja) string() string {
	switch f {
	case Espadas:
		return "Espadas"
	case Bastos:
		return "Bastos"
	case Oros:
		return "Oros"
	case Copas:
		return "Copas"
	}
	return "Invalido"
}
