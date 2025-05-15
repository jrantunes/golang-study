package main

import "fmt"

// a interface em go permite passar apenas assinaturas de métodos
type Pessoa interface {
	// Nome string // não pode adicionar propriedades
	Desativar() // qualquer struct que possuir este método estará implementando esta mesma interface
	// MetodoComRetorno() int  // este método deve retornar um inteiro
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado!\n", c.Nome)
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
	fmt.Printf("A empresa %s foi desativada!\n", e.Nome)
}

type Funcionario struct {
	Nome string
}

func (f Funcionario) Desativar(a int) {}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	cliente := Cliente{
		Nome:  "Claudio",
		Idade: 20,
		Ativo: true,
	}

	empresa := Empresa{
		Nome: "Natura",
	}

	// funcionario := Funcionario{
	// 	Nome: "Carlos",
	// }

	// ambas estão implementando a função Desativar da interface pessoa
	// ex: se a assinatura da interface possuisse parametros daria erro
	Desativacao(cliente)
	Desativacao(empresa)
	// e para essa implementação funcionar a assinatura da função Desativar de ambas as structs deve ser identica a da interface
	// Desativacao(funcionario) // erro (o método Desativar de funcionario não foi implementado da mesma forma da interface Pessoa)
}
