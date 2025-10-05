package main

import "fmt"

func main() {
	fmt.Println(soma(5, 5, 5))       // 15
	fmt.Println(soma(5, 5, 5, 5, 5)) // 25
}

// funções variádicas -- quando não sabemos a quantidade de parametros que vamos trabalhar
// ex: parametros de configuração de um banco de dados (esses parametros podem ser apenas um ou varios)
func soma(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}
