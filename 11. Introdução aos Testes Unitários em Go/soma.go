package main

import "fmt"

func main() {
	soma := soma(1, 2, 3)
	fmt.Println(soma)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
