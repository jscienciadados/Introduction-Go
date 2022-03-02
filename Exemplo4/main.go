package main

import "fmt"

// Cria uma struct
type pessoa struct {
	nome  string
	idade int
}

// Função para nova pessoa
func novaPessoa(nome string) *pessoa {

	p := pessoa{nome: nome}
	p.idade = 42

	return &p
}

// Função main
func main() {

	fmt.Println(pessoa{"Barbara", 29})

	fmt.Println(pessoa{nome: "Deborah Maria", idade: 17})

	fmt.Println(pessoa{nome: "Fabiane"})

	fmt.Println(&pessoa{nome: "Ana", idade: 28})

	fmt.Println(novaPessoa("Marina"))

	s := pessoa{nome: "Joao Souza", idade: 48}
	fmt.Println(s.nome)

	sp := &s
	fmt.Println(sp.idade)

	sp.idade = 58
	fmt.Println(sp.idade)
}