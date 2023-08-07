package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2) // canal com buffer
	
	canal <- "Olá mundo"
	canal <- "Programando em Go!"


	mensagem1 := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)

	// canal com buffer só vai bloquear tráfego quando atingir limite
}
