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

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Ímpar")
		}
	}
}

