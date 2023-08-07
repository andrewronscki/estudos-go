package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	
	cenariosDeTeste := []cenarioDeTeste {
		{"Avenida Paulista", "Avenida"},
		{"Rua Dos Bobos", "Rua"},
		{"Rodovia Bandeirantes", "Rodovia"},
		{"Praca das Rosas Paulista", "Tipo inválido"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
	} 

	for _, cenario := range cenariosDeTeste {
		recebido := TipoDeEndereco(cenario.enderecoInserido)
		
		if recebido != cenario.enderecoEsperado {
			t.Errorf(
				"O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				cenario.enderecoEsperado,
				recebido,	
			)
		}
	}
}