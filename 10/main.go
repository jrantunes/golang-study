package main

import "fmt"

// closures // funções anonimas
// dentro de uma função podemos ter outra função
// podem ser usadas para rodar algumas rotinas sem precisar armazenar um retorno pra elas
func main() {
	totalDobrado := func() int {
		return soma(2, 2, 2) * 2
	}()

	fmt.Println(totalDobrado) // 12
}

func soma(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}
