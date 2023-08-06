package main

import (
	"fmt"
)


func main() {
	usuario := map[string]string {
		"nome": "André",
		"sobrenome": "Wronscki",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
}