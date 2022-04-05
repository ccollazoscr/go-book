package cubo

/*Función que inicia con minúscula es privada*/
func eleva3(x float64) float64 {
	return x * x * x
}

/*Función que inicia con Mayúscula es publica*/
func Volumen(lado float64) float64 {
	return eleva3(lado)
}
