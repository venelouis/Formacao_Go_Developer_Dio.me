// Loop Repetição Hierárquica

// Escreva um programa que mostre no terminal todos os segundos do dia no formato de hora-minuto-segundo
// e que pare quando chegar no horário de 23:59:59 e depois mostre no terminal a mensagem "Fim!".

package main

import "fmt"

func main() {
	for h := 0; h <= 23; h++ {
		for m := 0; m <= 59; m++ {
			for s := 0; s <= 59; s++ {
				fmt.Printf("%d:%d:%d\n", h, m, s)
			}
		}
	}
	fmt.Println("Fim!")
}
