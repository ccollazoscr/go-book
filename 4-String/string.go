package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str1 := "hola mundo"
	fmt.Printf("%q mide %v\n", str1, len(str1))

	str2 := "Hola 物"
	fmt.Printf("%q mide %v\n", str2, len(str2))

	//Longitud en runtas:
	fmt.Printf("%q tine %v runas\n", str2, utf8.RuneCountInString(str2))

	//De string a porción
	cad := "una cadena"

	bytes := []byte(cad)
	runas := []rune(cad)

	//modificando el contenido de las porciones
	bytes[0] = 'U'
	runas = append(runas, '!')

	//Convirtiendo la porción a Cadena
	str3 := string(bytes)
	str4 := string(runas)

	fmt.Println(str3)
	fmt.Println(str4)

	showStringBuilderTest()
	textWithFormat()
}

func textWithFormat() {
	val := 3.1
	str := fmt.Sprintf("val = %.5f\n", val)
	fmt.Println(str)
}

func showStringBuilderTest() {
	var sb strings.Builder
	sb.WriteString("strings.Builder\n")
	l := sb.Len()
	for i := 0; i < l; i++ {
		sb.WriteRune('=')
	}

	sb.WriteString("\nEstamos escribiendo en un string builder")
	sb.WriteString("\nEsta es otra forma de escribir strings largos")

	fmt.Println(sb.String())
}
