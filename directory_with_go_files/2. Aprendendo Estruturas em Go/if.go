package main

import "fmt"

func main() {
	// SE o número é PAR ou ÍMPAR
	// resto = %
	/*
		se i%2 == 0
			imprimir: número par
		se não
			imprimir: número ímpar
	*/

	for i := 1; i <= 10; i++ { //i++ é a mesma coisa que i = i + 1
		if i%2 == 0 { //se o resto da divisão de i por 2 for igual a 0
			fmt.Println("Par") //imprime "Par"
		} else { //senão
			fmt.Println("Ímpar") //imprime "Ímpar"
		}
	}
}
