// go não tem while, tem apenas o for.
// utilizando o for como while:

package main

import "fmt"

func main() {
	x := 0        //inicialização
	for x <= 10 { //condição
		fmt.Println(x) // pós - saída
		x++            //execução (incremento)
	}
}
