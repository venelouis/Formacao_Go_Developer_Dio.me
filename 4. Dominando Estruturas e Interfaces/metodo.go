/*
método é uma função associada a um tipo particular.
Isto é, em POO(Prog. Orientada a Objetos), objeto é um valor(variável)
e o método é uma função associada a esse objeto.
*/
package main

type retangulo struct { //struct é um tipo de dado
	comprimento, altura int //atributos
}

// Este método "área" possui um tipo "retangulo" como receptor.
func (r retangulo) area() int { //método é uma função associada a um tipo particular (retangulo) e não a uma variável
	return r.comprimento * r.altura //retangulo é o tipo de dado e não uma variável (r) que é o receptor
}

// Métodos podem ser definidos por qualquer tipo de receptor,
// ponteiro ou valor. Aqui temos um exemplo do receptor de valor.
func (r *retangulo) perimetro() int { //método é uma função associada a um tipo particular (retangulo) e não a uma variável
	return 2*r.comprimento + 2*r.altura //retangulo é o tipo de dado e não uma variável (r) que é o receptor
}

func main() { //método é uma função associada a um tipo particular (retangulo) e não a uma variável
	r := retangulo{10, 5} //retangulo é o tipo de dado e não uma variável (r) que é o receptor

	// Aqui chamamos os 2 métodos definidos para o struct retangulo.
	println("area: ", r.area())           //método é uma função associada a um tipo particular (retangulo) e não a uma variável
	println("perimetro: ", r.perimetro()) //método é uma função associada a um tipo particular (retangulo) e não a uma variável

	// Go automaticamente trata a conversão entre valores e ponteiros
	// para métodos chamados. Você pode usar ponteiros para valores
	// ou valores em chamadas de métodos.
	rp := &r                               //retangulo é o tipo de dado e não uma variável (r) que é o receptor e o ponteiro (&) é o receptor de valor
	println("area: ", rp.area())           //método é uma função associada a um tipo particular (retangulo) e não a uma variável
	println("perimetro: ", rp.perimetro()) //método é uma função associada a um tipo particular (retangulo) e não a uma variável

	// Go escolhe automaticamente entre valores e ponteiros como
	// receptores. Aqui temos um exemplo de um receptor de ponteiro.
	// Aqui temos um exemplo de um receptor de ponteiro.

}
