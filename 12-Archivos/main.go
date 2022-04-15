package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	n, err := os.Stdout.Write([]byte("hola!\n"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Escritos:", n, "bytes")

	fmt.Println("Escribe 5 caracteres: ")
	datos := make([]byte, 5)
	n, err = os.Stdin.Read(datos)
	if err != nil {
		fmt.Println("Error leyendo:", err)
	}
	fmt.Println("Leidos:", n, "bytes", string(datos))

	EscribirEnArchivo()
	LeyendoArchivo()
	UsandoBufio()
}

func UsandoBufio() {
	buff := bytes.NewBufferString(`Hola que tal
	Probando texto
	multilinea`)
	sc := bufio.NewReader(buff)

	leido, err := sc.ReadString('\n')
	for err == nil {
		fmt.Print("Leida línea: ", leido)
		leido, err = sc.ReadString('\n')
	}

	if err == io.EOF {
		fmt.Print("Línea final: ", leido)
	} else {
		fmt.Print("Error inesperado:", err)
	}
}

func EscribirEnArchivo() {
	fmt.Printf("\nEscribiendo en archivo\n")
	fichero, err := os.Create("ejemplo.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fichero.Close()
	if _, err = fichero.Write([]byte("hola!")); err != nil {
		fmt.Println(err)
	}
}

func LeyendoArchivo() {
	fichero, err := os.Open("ejemplo.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer fichero.Close()
	leido := make([]byte, 256)
	n, err := fichero.Read(leido)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("leidos %v bytes: %s\n", n, string(leido[:n]))
}
