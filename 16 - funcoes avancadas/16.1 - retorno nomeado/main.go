package main

import (
	"fmt"
)

func calculosMatematicos(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	sum, sub := calculosMatematicos(10, 5)
	fmt.Println(sum, sub)
}