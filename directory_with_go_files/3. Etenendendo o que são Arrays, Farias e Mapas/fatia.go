// fatia = slice em inglês
// fatia é um pedaço de um array

// var fatia []float64
//fatia := make([]float64, 5)
//fatia:= array[0:5]

// Exemplo:

package main

import "fmt"

func main() {
	array := [7]float64{1, 2, 3, 4, 5, 6, 7}
	fatia := array[2:5]               // fatia recebe o array da posição 2 até a posição 5
	fatia2 := append(fatia, 8, 9, 10) // fatia2 recebe a fatia + 8, 9, 10
	fmt.Println(fatia, fatia2)        // printa a fatia
}

// Resultado: [3 4 5] [3 4 5 8 9 10]
