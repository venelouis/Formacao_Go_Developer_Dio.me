package main

import "fmt"

var a int = 20        //variável global (pode ser acessada por qualquer função) e não precisa ser usada (não dá erro)
var b string = "Nome" //variável global (pode ser acessada por qualquer função) e não precisa ser usada (não dá erro)

func main() {

	var c float64 = float64(a)                                  //tipo(variável) = tipo_desejado(variável) - conversão de tipos (casting) - precisa ser usada
	fmt.Printf("O valor de C é: %g e o valor de B é: %s", c, b) //%g para float64 e %s para string - não precisa ser usada (não dá erro)

}
