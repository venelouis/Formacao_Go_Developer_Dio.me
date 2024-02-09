// canais é uma forma de comunicação entre goroutines
// canais são um tipo de dado em Go
// canais são utilizados para enviar e receber dados entre goroutines

package main

import (
	"fmt"
	"time"
)

func ping(c chan string) { // função ping que recebe um canal c do tipo string
	for { // loop infinito
		c <- "ping" // envia a string "ping" para o canal c
	}
}
func imprimir(c chan string) { // função imprimir que recebe um canal c do tipo string
	for { // loop infinito
		msg := <-c                  // recebe a string do canal c e armazena na variável msg
		fmt.Println(msg)            // imprime a variável msg
		time.Sleep(time.Second * 1) // espera 1 segundo
	}
}
func main() {
	var c chan string = make(chan string) // cria um canal c do tipo string
	go ping(c)                            // executa a função ping em uma goroutine
	imprimir(c)                           // executa a função imprimir

	var entrada string
	fmt.Scanln(&entrada)
}
