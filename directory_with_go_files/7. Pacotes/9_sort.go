// ordenar dados arbitrários
// usaremos a função sort.Interface
// sort.Interface irá exigir 3 métodos: Len, Less e Swap

package main

import (
	"fmt"
	"sort"
)

type Dados struct { // struct para armazenar dados
	nome  string // nome da pessoa
	idade int    // idade da pessoa
}

type ParaNome []Dados // slice de Dados

func (ps ParaNome) Len() int { // método Len para sort.Interface que retorna o tamanho do slice de Dados (ParaNome)
	return len(ps) // retorna o tamanho do slice de Dados (ParaNome)
}

func (ps ParaNome) Less(i, j int) bool { // método Less para sort.Interface que retorna um booleano se o nome de i é menor que o nome de j no slice de Dados (ParaNome)
	return ps[i].nome < ps[j].nome // retorna verdadeiro se o nome de i for menor que o nome de j
}

func (ps ParaNome) Swap(i, j int) { // método Swap para sort.Interface que troca os valores de i e j no slice de Dados (ParaNome)
	ps[i], ps[j] = ps[j], ps[i] // troca os valores de i e j
}

func main() {
	criancas := []Dados{ // slice de Dados (criancas) com dados de crianças
		{"João", 10},
		{"Maria", 9},
		{"Pedro", 11},
		{"Ana", 8},
		{"Carlos", 12},
	}
	sort.Sort(ParaNome(criancas)) // ordena os dados de crianças pelo nome em ordem alfabética
	fmt.Println(criancas)         // imprime os dados de crianças ordenados pelo nome em ordem alfabética
}
