package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero2 *int) {
	*numero2 = *numero2 * -1
}


func main() {
	numero1 := 20
	numeroInvertido1 := inverterSinal(numero1)
	fmt.Println(numeroInvertido1)
	fmt.Println(numero1)
	// está enviando uma cópia do numero
	// não altera o valor do numero1

	numero2 := 20
	inverterSinalComPonteiro(&numero2)
	fmt.Println(numero2)
	// está enviaindo o ponteiro para a função
	// numero2 é alterado
}