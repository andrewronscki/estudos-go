package main

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		t.Parallel()

		ret := Retangulo{ 10, 12}
		esperado := float64(120)
		recebido := ret.area()

		if esperado != recebido {
			t.Fatalf(
				"Área recebida %f é direfente da esperada %f", 
				recebido, 
				esperado,
			)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		t.Parallel()
		
		circ := Circulo{ 10 }
		esperado := float64(math.Pi * 100)
		recebido := circ.area()

		if esperado != recebido {
			t.Fatalf(
				"Área recebida %f é direfente da esperada %f", 
				recebido, 
				esperado,
			)
		}
	})
}