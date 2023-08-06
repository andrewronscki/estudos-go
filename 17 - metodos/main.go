package main

import "fmt"

type Usuario struct {
	nome  string
	idade uint8
}

func (u Usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário '%s' no banco de dados\n", u.nome)
}

func (u Usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *Usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := Usuario{"Usuario 1", 20}
	usuario1.salvar()

	usuario2 := Usuario{"Usuario 2", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}