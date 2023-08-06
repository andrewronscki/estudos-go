package main

import "fmt"



func main() {
	// Aritiméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10/4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// Atribuição
	var str1 string = "String"
	str2 := "String"
	fmt.Println(str1, str2)

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// Operadores unários
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	
	numero--
	fmt.Println(numero)
	numero -= 15
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	numero /= 10
	fmt.Println(numero)

	numero %= 3
	fmt.Println(numero)

	// Não existe ternário no Go
	
	// ternário => texto := numero > 5 ? "Maior que 5" : "Menor que 5"
	// refatorando o ternário pra Go:
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}