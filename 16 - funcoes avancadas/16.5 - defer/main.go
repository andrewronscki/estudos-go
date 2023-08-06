package main

import "fmt"

func funcao1()  {
	fmt.Println("Executando funcao 1")
}

func funcao2()  {
	fmt.Println("Executando funcao 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado será retornado")
	fmt.Println("Verificando se aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao1() // defer -> adiar a executacao
	funcao2()
	aprovado := alunoAprovado(8, 7)
	fmt.Println(aprovado)
}