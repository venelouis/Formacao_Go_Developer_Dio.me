package main //pacote principal

import "fmt" //pacote para formatar a saída de dados na tela

func main() { /*função principal*/

	fmt.Println("Hello Word - Olá Mundo!") //imprime na tela e pula uma linha
	fmt.Print("Olá Mundo - Hello Word!")   //imprime na tela sem pular uma linha
	fmt.Println("Hello Word - Olá Mundo!") //imprime na tela e pula uma linha

}

// Para executar o código, basta digitar no terminal: go run hello.go
// Para compilar o código, basta digitar no terminal: go build hello.go
// Para executar o código compilado, basta digitar no terminal: ./hello
// Para criar um módulo, basta digitar no terminal: go mod init nome-do-modulo
// Para instalar um módulo, basta digitar no terminal: go get nome-do-modulo
// Para atualizar um módulo, basta digitar no terminal: go get -u nome-do-modulo
// Para remover um módulo, basta digitar no terminal: go mod tidy
// Para listar os módulos, basta digitar no terminal: go list -m all
// Para listar os módulos, basta digitar no terminal: go list -m -u all
// Para listar os módulos, basta digitar no terminal: go list -m -versions all
// Para listar os módulos, basta digitar no terminal: go list -m -versions -u all
