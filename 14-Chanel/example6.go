package main

import (
	"fmt"
	"time"
)

/*Demultiplexión con select*/
func main() {
	sms := make(chan string, 5)
	email := make(chan string, 5)
	carta := make(chan string, 5)

	go CentralMensajeria(sms, email, carta)

	sms <- "98798797987"
	email <- "coll3377897@gmail.com"
	carta <- "Banco central"
	email <- "coll8777@gmail.com"

	time.Sleep(time.Second)
}

func CentralMensajeria(sms, email, carta <-chan string) {
	for {
		select {
		case num := <-sms:
			fmt.Println("recibido SMS del número", num)
		case dir := <-email:
			fmt.Println("recibido email de dirección", dir)
		case rem := <-carta:
			fmt.Println("recibida carta de remitente", rem)
		}
	}
}
