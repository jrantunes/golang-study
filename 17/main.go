package main

import "fmt"

type Cliente struct {
	Nome  string
	Ativo bool
}

type Conta struct {
	Valor int
}

// recebe somente uma cópia da struct original
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %v foi destivado!\nValor de ativo(dentro de Desativar()): %t\n", c.Nome, c.Ativo)
}

// manipula a struct original
func (c *Cliente) DesativarOriginal() {
	c.Ativo = false
	fmt.Printf("O cliente %v foi destivado!\nValor de ativo(dentro de DesativarOriginal()): %t\n", c.Nome, c.Ativo)
}

// exemplo em que faz sentido manipular apenas a cópia da struct e não a original
// a função abaixo não altera o valor original da struct
func (c Conta) simular(valor int) int {
	c.Valor += valor
	fmt.Printf("Valor simulado %d\n", c.Valor)
	return c.Valor
}

// a função abaixo altera o valor original da struct
// func (c *Conta) simular(valor int) int {
// 	c.Valor += valor
// 	fmt.Printf("Valor simulado %d\n", c.Valor)
// 	return c.Valor
// }

// metódo comum para criação de structs em go
func NewConta() *Conta {
	// qualquer alteração feita nesta struct será refletida em todos os lugares que utilizam a mesma
	return &Conta{Valor: 0}
}

// EX
func NewCliente(nome string) *Cliente {
	return &Cliente{
		Nome:  nome,
		Ativo: true,
	}
}

func DesativarCliente(c *Cliente) {
	c.Ativo = false
}

func main() {
	c := Cliente{
		Nome:  "Teste",
		Ativo: true,
	}
	c.Desativar()
	fmt.Printf("Valor de ativo: %t\n", c.Ativo) // true
	c.DesativarOriginal()
	fmt.Printf("Valor de ativo: %t\n", c.Ativo) // false

	conta := Conta{
		Valor: 100,
	}
	conta.simular(200) // printa 300
	fmt.Printf("Valor na conta %d\n", conta.Valor)

	// cliente com metódo construtor
	cliente := NewCliente("Claudio")
	fmt.Printf("Nome cliente: %v\nAtivo: %t\n\n", cliente.Nome, cliente.Ativo) // Claudio true
	DesativarCliente(cliente)
	fmt.Printf("Nome cliente: %v\nAtivo: %t\n\n", cliente.Nome, cliente.Ativo) // Claudio false
}
