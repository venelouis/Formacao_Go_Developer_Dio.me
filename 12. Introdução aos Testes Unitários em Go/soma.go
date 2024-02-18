package main

import "fmt"

func main() {
	soma := soma(1, 2, 3)
	fmt.Println(soma)
}

func soma(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}
