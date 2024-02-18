// a crypto são exemplos criptografados de hash
// a hash são exemplos não criptografados de hash

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()         // cria uma nova hash
	h.Write([]byte("test")) // escreve os dados na hash
	bs := h.Sum(nil)        // pega o valor da hash
	fmt.Println(bs)         // imprime o valor da hash
}
