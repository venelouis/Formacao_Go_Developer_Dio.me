// Goroutine é uma função que é executada de forma concorrente com outras funções.
// A função main é executada de forma concorrente com a função f através da palavra-chave go.

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
} 