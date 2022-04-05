package main

type ContadorApuntadores int

func (c *ContadorApuntadores) Incrementa() {
	*c++
}

func (c *ContadorApuntadores) Reinicia(nuevoValor int) {
	*c = ContadorApuntadores(nuevoValor)
}
