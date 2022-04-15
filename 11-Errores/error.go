package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	division, err := Divide(6, 3)
	if err != nil {
		fmt.Println("Ha ocurrido un error", err)
	} else {
		fmt.Println("La división es", division)
	}

	division, err = Divide(6, 0)
	if err != nil {
		fmt.Println("Ha ocurrido un error. Err:", err)
	} else {
		fmt.Println("La división es", division)
	}

	_, err2 := RaizCuadrada(-1)
	if err2 != nil {
		fmt.Println("Ha ocurrido un error. Err:", err2)
	}

	//********************************************
	fmt.Printf("\nErrores Centinela\n")
	s := NewAlmacen(30)
	err = s.Guarda("un_numero", 12345)
	switch err {
	case ErrYaExiste:
		fmt.Println(err, ": Prueba con otra clave única")
	case ErrNoCapacidad:
		fmt.Println(err, ": No se pueden guardar más elementos")
	case ErrClaveInvalida:
		fmt.Println(err, ": prueba con otra clave bien formateada")
	case nil:
		fmt.Println("operación llevada a cabo con éxito!")
	default:
		fmt.Println("error desconocido:", err)
	}

	//********************************************
	fmt.Printf("\nErrores com parámetro\n")
	dog := Perro{TieneHambre: true}
	alimentos := []Comida{
		ComidaPajaro, ComidaGato, ComidaPerro, ComidaHumano,
	}

	for _, f := range alimentos {
		err = dog.Alimenta(f)
		switch err.(type) {
		case NoApetecible:
			fmt.Println(err, "-> prueba otro tipo de comida")
		case NoHambre:
			fmt.Println(err, "-> espera unas horas")
		case error:
			fmt.Println(err, "(no esperaba esto)")
		}
	}

	//********************************************
	fmt.Printf("\nEnvolviendo errores\n")
	err3 := ErrorTv{}
	fmt.Println(errors.Unwrap(err3))

	err3 = ErrorTv{Causa: errors.New("Se perdio la señal")}
	fmt.Println(errors.Unwrap(err3))

	err4 := ErrorTv{Causa: ErrorComponente{Nombre: "Condensador"}}
	var errTV ErrorTv
	if errors.As(err4, &errTV) { //Si encuentra el error envuelto, lo asigna la la variable errTV y devuelve true
		fmt.Println("Encontrado en cadena de error:", errTV)
	}

	//*********************************************
	fmt.Printf("\nUso de función defer\n")
	CambiarRegistro()

	//*********************************************
	fmt.Printf("\nUso de función recover para recuperarse de panico\n")
	fmt.Println("Invocando una función")
	funcion1()
	fmt.Println("Saliendo con normalidad")
}

func Divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func RaizCuadrada(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("%f: no hay raiz cuadrada real", n)
	}
	return math.Sqrt(n), nil
}
