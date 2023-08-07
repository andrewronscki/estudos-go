package main

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func EscreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f\n",f.area())
}

func main() {
	r := Retangulo{10, 15}
	EscreverArea(r)

	c := Circulo{10}
	EscreverArea(c)
}