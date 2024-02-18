package main //package main

import "fmt" //importa o pacote fmt

func main() { //função principal
	x := 4                //variável x recebe 4
	if x == 2 || x == 3 { //se x for igual a 2 ou x for igual a 3
		fmt.Print("Sim, x é 2 OU 3") //imprime "Sim, x é 2 OU 3"
	} else { //senão
		fmt.Print("X não é nem 2 nem 3!") //imprime "X não é nem 2 nem 3!"
	}
}
