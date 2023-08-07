package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia dos imigrantes")
	fmt.Println(tipoEndereco)
}