// Exemplo de loop infinito

package main

import "fmt"

func main() {

	i := 1
	for i < 2 {
		fmt.Println("loop")
		break
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 { // se o resto da divisão de i por 2 for igual a 0
			continue
		}
		fmt.Println(i)
	}
}

// Sem o break o loop seria infinito

// Exemplo utilizando continue
