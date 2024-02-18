// ponteiro(ptr) é um tipo especial de variável que armazena o endereço de memória de outra variável.
// armazena um valor na memória, mas não é o valor propriamente escrito.
// para criar um ponteiro, usamos o operador "&" antes da variável que queremos referenciar.
// para criar um ponteiro, usamos o operador "*" antes do tipo de variável que queremos referenciar.
// para desreferenciar um ponteiro, usamos o operador "*" antes do ponteiro.
// para criar um ponteiro usamos a função new() que retorna um ponteiro.

package main

import "fmt"

func inicial(xPtr *int) {
	*xPtr = 3
}

func main() {
	x := 5
	inicial(&x)
	fmt.Println(x)
}
