// aceita um conjunto de dados e o reduz a um tamanho fixo menor
// hashes são usadas em TUDO na programação, principalmente para buscar DADOS e DETECTAR
// em Go são dividias em: CRIPTOGRAFADAS e NÃO CRIPTOGRAFADAS
// lista das NÃO CRIP: adler32, crc32, crc64, fnv
// lista das CRIP: md5, sha1, sha256, sha512
// nesse código vamos criar uma HASH e escrever nossos dados nele.
// alternativa ao sha1-online.com que está como "inseguro" (inclusive travou meu pc): https://md5hashing.net/

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()    // cria uma nova hash
	h.Write([]byte("test")) // escreve os dados na hash
	v := h.Sum32()          // pega o valor da hash
	fmt.Println(v)          // imprime o valor da hash
}
