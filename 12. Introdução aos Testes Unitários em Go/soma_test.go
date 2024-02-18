package main

import "testing" // Importando o pacote de testes, mas para ele rodar é necessário go.mod que é criado com: go mod init + soma_test.go
// para rodar todos os testes utilize o comando no terminal: go test -v

// Teste unitário para a função soma
// padrão triple A - AAA
// Arrange - Organizar (preparar o ambiente para o teste - preparar)
// Act - Agir (executar a função que será testada - rodar)
// Assert - Afirmar (verificar as asserções)
func TestSoma(t *testing.T) {
	// Arrange
	teste := soma(1, 2, 3)
	resultado := 6
	// Act
	if teste != 6 {
		// Assert
		t.Errorf("Resultado da soma é inválido: Resultado: %d. Esperado: %d", teste, resultado)
	} else {
		t.Logf("Resultado da soma é válido: Resultado: %d. Esperado: %d", teste, resultado)
	}
}

// exemplo de erro:
func TestSomaErrada(t *testing.T) {
	// Arrange
	teste := soma(1, 2, 3)
	resultado := 6
	// Act
	if teste == resultado {
		// Assert
		t.Errorf("Resultado da soma é inválido: Resultado: %d. Esperado diferente de: %d", teste, resultado)
	} else {
		t.Logf("Resultado da soma é válido: Resultado: %d. Esperado: diferente de %d", teste, resultado)
	}
}
