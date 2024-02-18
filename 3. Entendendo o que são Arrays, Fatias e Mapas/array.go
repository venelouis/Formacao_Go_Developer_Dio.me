package main

import "fmt"

func main() {
	var array [5]float64 // array de 5 posições do tipo float64
	array[0] = 7.4       // atribuindo valor a posição 0 do array
	array[1] = 6.5       // atribuindo valor a posição 1 do array
	array[2] = 5.5
	array[3] = 8.8
	array[4] = 9.9

	var total float64 = 0             // variável total do tipo float64
	for i := 0; i < len(array); i++ { // for para percorrer o array e somar os valores das posições
		total += array[i] // total recebe total + array na posição i
	}

	fmt.Println(total / float64(len(array))) // printa o total dividido pelo tamanho do array (len(array)) convertido para float64
	// 7.62 é o resultado da soma dos valores das posições do array dividido pelo tamanho do array
}

// A diferença entre Array 32 e 64 bits é que o 32 bits é menor e o 64 bits é maior.

// O Array é uma estrutura de dados que armazena uma coleção de elementos do mesmo tipo.
// O tamanho do Array é fixo e é definido na declaração do Array.
// O tamanho do Array é parte do seu tipo, então [5] int e [25] int são tipos diferentes.
