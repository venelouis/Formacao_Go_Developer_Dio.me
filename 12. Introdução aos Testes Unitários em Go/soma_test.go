package main

import "testing" // Importando o pacote de testes, mas para ele rodar é necessário go.mod que é criado com: go mod init + soma_test.go

func TestSoma(t *testing.T) {
	total := soma(1, 2, 3)
	if total != 6 {
		t.Errorf("Resultado da soma é inválido: Resultado: %d. Esperado: %d", total, 6)
	} else {
		t.Logf("Resultado da soma é válido: Resultado: %d. Esperado: %d", total, 6)
	}
}

func TestSomaErrada(t *testing.T) {
	total := soma(1, 2, 3)
	if total == 6 {
		t.Errorf("Resultado da soma é inválido: Resultado: %d. Esperado diferente de: %d", total, 0)
	} else {
		t.Logf("Resultado da soma é válido: Resultado: %d. Esperado: diferente de %d", total, 6)
	}
}
