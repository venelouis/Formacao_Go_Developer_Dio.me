// Select é uma forma de concorrência que permite que você execute várias operações ao mesmo tempo e escolha a primeira que terminar.
// Funciona como um switch, mas para canais.
// Permitindo que você execute várias operações ao mesmo tempo e escolha a primeira que terminar.

package main

// Cada canal receberá um valor após algum tempo, para simular operações concorrentes.
// por exemplo o bloqueio  de operações RPC ou HTTP.
// Cada canal é um canal de comunicação entre goroutines.
// Canais são um tipo de dado em Go.
// Canais são utilizados para enviar e receber dados entre goroutines.
import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string) // cria um canal c1 do tipo string
	c2 := make(chan string) // cria um canal c2 do tipo string

	go func() { // executa a função anônima em uma goroutine
		time.Sleep(time.Second * 1) // espera 1 segundo
		c1 <- "um"                  // envia a string "um" para o canal c1
	}()
	go func() { // executa a função anônima em uma goroutine
		time.Sleep(time.Second * 2) // espera 2 segundos
		c2 <- "dois"                // envia a string "dois" para o canal c2
	}()

	for i := 0; i < 2; i++ { // loop de 0 a 1
		select { // seleciona o primeiro canal que estiver pronto
		case msg1 := <-c1: // recebe a string do canal c1 e armazena na variável msg1
			fmt.Println("recebido", msg1) // imprime a string "recebido" e a variável msg1
		case msg2 := <-c2: // recebe a string do canal c2 e armazena na variável msg2
			fmt.Println("recebido", msg2) // imprime a string "recebido" e a variável msg2
		}
	}
}
