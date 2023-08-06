package main

import (
	"errors"
	"fmt"
)

func main() {
	// Todo tipo de dado possui seu valor inicial

	// int, int8, int16, int32, int64
	// Por inferência ele vai se basear na arquitetura do sistema operacional
	var numero1 int64 = -10000000
	fmt.Println(numero1)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// float32, float64
	var numero5 float32 = 123.45
	fmt.Println(numero5)

	var numero6 float64 = 123456789.23
	fmt.Println(numero6)


	// string
	var str1 string = "Texto"
	fmt.Println(str1)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	// número que ele está na tabela ascii
	fmt.Println(char)

	// bool
	var booleano bool
	fmt.Println(booleano)

	// error
	var erro1 error
	fmt.Println(erro1)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}