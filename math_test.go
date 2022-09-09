package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSub(t *testing.T) {
	total := Sub(15, 15)

	if total != 0 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado: %d", total, 0)
	}
}

func TestTimes(t *testing.T) {
	total := Times(1, 15)

	if total != 15 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado: %d", total, 15)
	}
}
