// declarar meu pacote principal
package main

//importar a função fmt
import "fmt"

//declaração da variável CONST do ponto de ebulição da água em Fahrenheit
const ebulicaoFahrenheit = 212.0

//função principal
func main() {

	var tempFahrenheit = ebulicaoFahrenheit         //variável do valor da temperatura em graus Fahrenheit
	var tempCelsius = (tempFahrenheit - 32) * 5 / 9 //Conversão de F (Fahrenheit) para C (Celsius)
	//Apareça na tela o resultado
	/*
		fmt.Println("Atemperatura de ebulição da água em Fahrenheit é = ", tempFahrenheit)
		fmt.Println("A temperatura de ebluição da água em Celsius é = ", tempCelsius)
	*/
	// também pode ser escrito de outra maneira com Printf e vai ficar:
	fmt.Printf("A temperatura de ebulição da água em Fahrenheit é = %g e a temperatura da água em Celsius é = %g, %T", tempFahrenheit, tempCelsius, tempCelsius)
	//*%g (porque é um ponto flutuante)
	// A gente espera que apareça F = 212 e C = 100

}

