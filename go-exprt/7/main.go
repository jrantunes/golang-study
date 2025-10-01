package main

import "fmt"

func main() {
	// maps - hashtables (estrutura de dados com chave e valor)
	// criação de um map
	salarios := map[string]int{"Juninho": 1000, "Pedrinho": 2000, "Claudinho": 3000}
	fmt.Println(salarios["Juninho"]) // 1000

	// para remover um dado no map
	delete(salarios, "Juninho")
	fmt.Println(salarios["Juninho"]) // 0

	// para adicionar um dado no map
	salarios["Pepinho"] = 4000
	fmt.Println(salarios["Pepinho"]) // 4000

	// criando um map vazio com make
	// salariosmake := make(map[string]int)
	// criando um map vazio sem make
	// salariosvazio := map[string]int{}

	// percorrendo um map
	for nome, salario := range salarios {
		fmt.Printf("Nome: %s; Salário: %d\n", nome, salario)
	}

	// percorrendo e ignorando um valor
	// _ blank identifier
	for _, salario := range salarios {
		fmt.Printf("Salário: %d\n", salario)
	}
}
