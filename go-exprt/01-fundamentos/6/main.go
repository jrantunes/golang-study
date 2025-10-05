package main

import "fmt"

func main() {
	// criação de um slice
	// quando trabalhamos com slices estamos trabalhando tambem com arrays por debaixo dos panos
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// len = 9; cap = 9; valor = [1 2 3 4 5 6 7 8 9]
	fmt.Printf("len = %d; cap = %d; valor = %v \n", len(slice), cap(slice), slice)

	// [:index] ou [index:] cria um novo slice
	// [:n] pega todos os n valores a partir do indice 0
	// [n:] pega todos os valores do indice n em diante

	// com [:n] a capacidade não altera diferente de [n:] em que a capacidade é igual a capacidade original - n

	// len = 0; cap = 9; valor = []
	fmt.Printf("len = %d; cap = %d; valor = %v \n", len(slice[:0]), cap(slice[:0]), slice[:0])

	// len = 4; cap = 9; valor = [1 2 3 4]
	fmt.Printf("len = %d; cap = %d; valor = %v \n", len(slice[:4]), cap(slice[:4]), slice[:4])

	// len = 7; cap = 7; valor = [3 4 5 6 7 8 9]
	fmt.Printf("len = %d; cap = %d; valor = %v \n", len(slice[2:]), cap(slice[2:]), slice[2:])

	// append -> adiciona elementos novos no slice
	// por debaixo dos panos o append gera um novo array na memória com todos os elementos duplicando a capacidade inicial dele
	// nesse exemplo a capacidade inicial era de 9 e depois do append ela passa a ser 18
	slice = append(slice, 10)

	// len = 10; cap = 18; valor = [1 2 3 4 5 6 7 8 9 10]
	fmt.Printf("len = %d; cap = %d; valor = %v \n", len(slice), cap(slice), slice)

	// criando uma slice com make
	makeslice := make([]int, 1)
	makeslice[0] = 1

	fmt.Printf("%v\n", makeslice) // [1]
}
