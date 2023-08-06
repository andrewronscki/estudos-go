package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	// Go não possui classes, o struct é o que se mais assemelha

	var u usuario
	fmt.Println(u)
	u.nome = "André"
	u.idade = 28

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	u2 := usuario{"João", 28, enderecoExemplo}

	u3 := usuario{nome: "João"}

	fmt.Println(u, u2, u3)
}