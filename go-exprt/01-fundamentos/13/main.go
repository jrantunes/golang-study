package main

import "fmt"

// métodos em structs
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado!\n", c.Nome)
}

func main() {
	cliente := Cliente{
		Nome:  "Claudio",
		Idade: 20,
		Ativo: true,
	}

	cliente.Desativar()
}
