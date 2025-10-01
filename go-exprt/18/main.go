package main

import "fmt"

// interfaces vazias eram utilizadas em versões anteriores do go onde não haviam generics
// servem para termos liberdade nos tipos de dados

// a interface abaixo implementa todo mundo
// type x interface {}

// devemos utilizar interface vazia com bastante moderação

func main() {
	// exemplo de utilização de interfaces vazias
	var x interface{} = 10
	var y interface{} = "Hello world"

	showType(x) // O tipo da variável é int e o valor é 10
	showType(y) // O tipo da variável é string e o valor é Hello world
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}

// hoje em dia possuimos um recurso chamado de generics que consegue substituir bastante a necessidade de
// trabalhar com interfaces vazias
