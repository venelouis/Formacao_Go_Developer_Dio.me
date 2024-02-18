// defer: escalona as nossoas funções para serem executadas após a função que as contém
// defer tem como principal função adiar a execução de uma função até que a função que a contém retorne

package main

import "fmt"

func dia1() {
	fmt.Println("Domingo")
}
func dia2() {
	fmt.Println("Segunda-feira")
}

func main() {
	defer dia2() //defer é usado para adiar a execução de uma função até que a função que a contém retorne
	dia1()       //aqui a função dia1 é chamada
}
