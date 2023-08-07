package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enrecoMinusculo := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enrecoMinusculo, " ")[0]

	enderecoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo inválido"
}