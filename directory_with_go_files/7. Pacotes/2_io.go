// para saber mais sobre pacotes o site da documentação é: https://pkg.go.dev/
// este exemplo pode ser visto em: https://cs.opensource.google/go/go/+/go1.21.6:src/io/io.go;l=314 e executado tbm em: https://go.dev/play/p/pslFrsvNIuX

package main

import (
	"io"  // pacote de entrada e saída
	"log" // pacote de log
	"os"  // pacote de sistema operacional (entrada e saída de dados do sistema operacional)
)

func main() {
	if _, err := io.WriteString(os.Stdout, "Hello World"); err != nil {
		// escreve na saída padrão (os.Stdout) a string "Hello World"
		// se houver erro, imprime o erro no log e termina o programa
		log.Fatal(err) // imprime o erro no log e termina o programa
	}

}
