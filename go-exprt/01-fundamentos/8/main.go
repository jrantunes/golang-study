package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(soma(1, 2))        // 3 false
	fmt.Println(soma(5, 5))        // 10 true
	fmt.Println(somaComErro(5, 5)) // 0 A soma é maior ou igual a 10

	// armazenando o retorno da função com possibilidade de erro
	valor, err := somaComErro(5, 5)
	if err != nil {
		fmt.Println(err) // A soma é maior ou igual a 10
	}
	fmt.Println(valor) // 0
}

// declaração de função
// func nome(param1 tipo, param2 tipo, ...) tipo_do_retorno // deixar em branco se não houver retorno {}
// declaração de tipos aninhados -- func nome(param1, param2, param3 tipo) {} // param1, param2 e param3 são do mesmo tipo
// função com mais de um valor no retorno -- func nome() (tipo_de_retorno1, tipo_de_retorno2) {}
// funções com mais de um valor no retorno são utilizadas bastante em relação a erros (o go não possui exceptions (try, catch etc...))
func soma(a, b int) (int, bool) {
	if a+b >= 10 {
		return a + b, true
	}
	return a + b, false
}

// funções que podem retornar erros
func somaComErro(a, b int) (int, error) {
	if a+b >= 10 {
		return 0, errors.New("A soma é maior ou igual a 10")
	}
	return a + b, nil // sem erro / erro nulo
}
