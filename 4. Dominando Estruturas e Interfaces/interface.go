// Interfaces são coleções de métodos.

package main

import (
	"fmt"
	"math"
)

// Aqui temos uma interface usada para formas geométricas.
type geometria interface {
	area() float64
	perimetro() float64
}

/* Para o nosso exemplo, implementaremos essa interface em tipos quadrado e circulo.
Área de um quadrado: lado * lado ou lado²
Perímetro de um quadrado: 4 * lado = soma dos lados. */

type quadrado struct {
	lado float64
}
type circulo struct { //área do círculo: π(pi) * raio² perimetro do círculo: 2 * π(pi) * raio
	raio float64
}

// Para implementar uma interface em Go, precisamos implementar todos os métodos na interface.
// Aqui implementamos a interface geometria para o tipo quadrado.
func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

// Aqui implementamos a interface geometria para o tipo circulo.
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// Se uma variável tem um tipo de interface, podemos chamar métodos que estão na interface.
// Aqui temos uma função que usa a interface geometria.
func medida(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	q := quadrado{lado: 10}
	c := circulo{raio: 5}

	// Os tipos de círculo e quadrado implementam a interface geometria,
	// então podemos usar instâncias dessas estruturas como argumentos para medida.
	medida(q)
	medida(c)
}
