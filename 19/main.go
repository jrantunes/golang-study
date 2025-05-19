package main

import "fmt"

// type assertion é um mecanismo que permite recuperar o tipo concreto armazenado dentro de uma variável de interface.
// isso é útil quando você precisa acessar métodos ou propriedades específicas de um tipo concreto que não fazem parte da interface.

func main() {
	var minhaVar interface{} = "Hello world"
	println(minhaVar)          // (0x...,0x...)
	println(minhaVar.(string)) // Hello world
	// se a asserção falhar (ou seja, se a interface não contiver o tipo especificado), o programa entrará em pânico.
	// para evitar isso, podemos usar a forma segura com dois valores de retorno
	res, ok := minhaVar.(int)
	fmt.Printf("Valor de res é %v e valor de ok é %t", res, ok) // Valor de res é 0 e valor de ok é false
	// res, ok := minhaVar.(string)
	// fmt.Printf("Valor de res é %v e valor de ok é %t", res, ok) // Valor de res é Hello world e valor de ok é true
	// res := minhaVar.(int)
	// println(res) // panic: interface conversion: interface {} is string, not int
}
