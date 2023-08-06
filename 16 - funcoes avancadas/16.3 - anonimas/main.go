package main

import (
	"fmt"
)


func main() {
	func ()  {
		fmt.Println("Hello World")
	}()

	func (text string)  {
		fmt.Println(text)
	}("Parâmetro")

	retorno := func (texto string) string  {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Texto teste")

	fmt.Println(retorno)
}