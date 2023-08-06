package main

import "fmt"

var n int

func init() {
	fmt.Println("Execução função init")
	n = 10
}



func main() {
	fmt.Println("Execução função main")
	fmt.Println(n)
}