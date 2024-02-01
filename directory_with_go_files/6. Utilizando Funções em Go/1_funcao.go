// função também é chamada de procedimento ou sub-rotina
// função é um bloco de código que executa uma tarefa específica
// ele pega um dado de entrada e dá um dado de saída

package main

import (
	"fmt"
)

func media(lista []float64) float64 { //função que calcula a média
	total := 0.0
	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista)) //média
}

func main() { //programa que calcula a média de uma sala de aula
	lista := []float64{98, 93, 77, 82, 83} //lista de notas
	fmt.Println(media(lista))              //média
	/*
		total := 0.0
		for _, valor := range lista {
			total += valor
		}
		fmt.Println(total / float64(len(lista))) //média
	*/
}
