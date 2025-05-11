package main

import "fmt"

func main() {
	// array com 3 posições fixas do tipo int
	var arr [3]int

	// atribuindo valores no array
	arr[0] = 1
	arr[1] = 2
	arr[2] = 10
	// arr[3] = 20 // out of bounds (index invalido)

	// tamanho do array
	fmt.Println(len(arr))
	// ultimo elemento do array
	fmt.Println(arr[len(arr)-1])

	// percorrendo um array
	for index, value := range arr {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", index, value)
	}

	fmt.Printf("\nANTIGO\n\n")

	// laço de repetição modo antigo
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Valor em %d é %d\n", i, arr[i])
	}

	fmt.Printf("\nATUAL\n\n")

	// laço de repetição modo atual
	for i := range len(arr) {
		fmt.Printf("Valor em %d é %d\n", i, arr[i])
	}
}
