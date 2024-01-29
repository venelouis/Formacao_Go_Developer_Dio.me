// são coleções de "campos"
// para agrupar dados
// formar registros
// type() struct
package main
type pessoa struct{
	nome string
	idade int
}
import "fmt"

func main(){
	fmt.Println(pessoa{'Ana', 54})
	fmt.Println(pessoa{'João', 20})
}