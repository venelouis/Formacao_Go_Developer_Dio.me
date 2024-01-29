package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println(pessoa{"Ana", 54}) //tem que ser aspas duplas
	fmt.Println(pessoa{"João", 20})
}

// são coleções de "campos"
// para agrupar dados
// formar registros
// type() struct
