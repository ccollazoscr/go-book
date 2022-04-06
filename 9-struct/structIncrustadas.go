package main

type Aparato struct {
	Nombre string
	Precio int
}

func (g *Aparato) AplicaDecuento(d float32) {
	g.Precio = int(float32(g.Precio) * (1 - d))
}

type Telefono struct {
	Aparato
	Pulgadas int
	Bateria  int
}
