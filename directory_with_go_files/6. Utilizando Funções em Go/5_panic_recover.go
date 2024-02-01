// panic: erro do programador/ erro de execução
// panic: interrompe a execução do programa e imprime um erro
// recover: tenta recuperar o controle do programa (interrope o panic)

package main

import "fmt"

func main() {
	defer func() {
		x := recover()
		fmt.Println(x)
	}()
	panic("Pânico")
}

/*
func main() {
	panic("Pânico")
	x := recover()
	fmt.Println(x)
}
*/
