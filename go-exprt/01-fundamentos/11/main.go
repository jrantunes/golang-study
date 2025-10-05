package main

import "fmt"

// structs:  tipos de dados que permitem agrupar várias variáveis sob um único nome, funcionando como um contêiner para diferentes valores.
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	c := Cliente{
		Nome:  "Carlos",
		Idade: 20,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", c.Nome, c.Idade, c.Ativo)

	c.Ativo = false

	fmt.Println(c.Ativo)

	var c2 Cliente = Cliente{
		Nome:  "pedrinhi",
		Idade: 10,
		Ativo: true,
	}

	fmt.Println(c2.Nome)
}
