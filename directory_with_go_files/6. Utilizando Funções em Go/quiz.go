// O código abaixo apresenta a seguinte saída:

package main // pacote principal

import "fmt" // pacote para formatar a saída

func plus(a int, b int) int { // função que soma dois números inteiros

	return a + b // retorna a soma dos dois números

}

func plusPlus(a, b, c int) int { // função que soma três números inteiros

	return a + b + c // retorna a soma dos três números

}

func main() { // função principal

	res := plus(1, 2) // variável que armazena o resultado da função plus com os parâmetros 1 e 2

	fmt.Println("1+2 =", res) // imprime o resultado da soma dos dois números com o texto "1+2 =" e o valor da variável res

	res = plusPlus(1, 2, 3) // variável que armazena o resultado da função plusPlus com os parâmetros 1, 2 e 3

	fmt.Println("1+2+3 =", res) // imprime o resultado da soma dos três números com o texto "1+2+3 =" e o valor da variável res

}

// Explicando  e mostrando a saída do código que é:
// 1+2 = 3
// 1+2+3 = 6
// porque a função plus soma dois números inteiros e a função plusPlus soma três números inteiros.
// A função main chama as duas funções e imprime o resultado das somas.
