// closure é uma função que referencia variáveis de fora do seu corpo
// closure: função "chamar" uma variável que está em outra função
// capacidade de criar funções dentro de funções

package main

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	println(increment()) //1
	println(increment()) //2
	println(increment()) //3
}
