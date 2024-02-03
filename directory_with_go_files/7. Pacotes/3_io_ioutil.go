// o exemplo da aula foi descontinuado conforme o próprio link da documentação:
// https://pkg.go.dev/io/ioutil#example-WriteFile

// a função agora simplesmente se chama o.WtriteFile e pode ser vista em: https://pkg.go.dev/os#WriteFile
// o WriteFile é uma função que permite escrever dados em um arquivo.

package main

import (
	"log" // pacote de log
	"os"  // pacote de sistema operacional (entrada e saída de dados do sistema operacional)
)

func main() {
	err := os.WriteFile("testdata/hello", []byte("Hello, Gophers!"), 0666)
	// escreve no arquivo "testdata/hello" a string "Hello, Gophers!"
	//e permite que o arquivo seja editado por qualquer usuário
	if err != nil {
		// se houver erro, imprime o erro e encerra o programa
		// se não houver erro, imprime "Arquivo salvo com sucesso!"
		log.Fatal(err) // imprime o erro e encerra o programa
	}
}
