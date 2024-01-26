package main

import "fmt"

func main() {
	x := 4
	if x == 2 || x == 3 {
		fmt.Print("Sim, x é 2 OU 3")
	} else {
		fmt.Print("X não é nem 2 nem 3!")
	}
}

