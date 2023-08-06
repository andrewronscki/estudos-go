package main

import (
	"fmt"
)


func main() {
	// Array é inflexivel, possui tamanho fixo
	var array1 [5]int
	array1[0] = 1

	fmt.Println(array1)

	array2 := [5]int{1,2,3,4,5}
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5,6,7,8}
	fmt.Println(array3)

	// Slice é baseado no array, porém é flexivel, possui tamanho dinâmico
	slice1 := []int{1,2,3,4,5}
	fmt.Println(slice1)

	slice1 = append(slice1, 6)
	fmt.Println(slice1)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Arrays internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade
	
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	// Go gerencia o array interno alocado pro Slice
	// Ao estourar a capacidade ele cria outro array com length e capacidade novos
}