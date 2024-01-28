// coleção ordenada (lista) pares de chave e valor
// var x map [string] int
// x é um mapa de strings para inteiros

package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	x["key2"] = 20
	fmt.Println(x["key"], x["key2"])

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	elemento["Be"] = "Berílio"
	elemento["B"] = "Boro"
	elemento["C"] = "Carbono"
	elemento["N"] = "Nitrogênio"
	elemento["O"] = "Oxigênio"
	elemento["F"] = "Flúor"
	elemento["Ne"] = "Neônio"
	elemento["Na"] = "Sódio"
	elemento["Mg"] = "Magnésio"

	fmt.Println(elemento)
}

// Resultado: map[B:Boro Be:Berílio C:Carbono F:Flúor H:Hidrogênio He:Hélio Li:Lítio Mg:Magnésio N:Nitrogênio Na:Sódio Ne:Neônio O:Oxigênio]
// Resultado: 10 20
// Mapas são uma coleção de pares chave/valor sem ordem definida.
// Para criar um mapa vazio, use a função make:
// x := make(map[string]int)
// Para adicionar itens ao mapa, basta usar x [chave] = valor.
// Para recuperar um valor, use x [chave].
