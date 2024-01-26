package main

import "fmt"

func main() {
	//fmt.Println(`1,2,3...`) //imprime 1,2,3...

	/*
		i := 1
		for i <= 10 {
			fmt.Println(i)
			i = i + 1
		}
	*/

	for i := 1; i <= 10; i++ { //i++ é a mesma coisa que i = i + 1 (incrementa 1)
		fmt.Println(i) //imprime o valor de i
	}
}
