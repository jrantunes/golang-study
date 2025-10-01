package main

// importação de pacote
import "fmt"

// importação de múltiplos pacotes
/*
	import (
		"fmt"
		"net/http"
	)
*/

type ID int

var (
	a float64 = 1.2
	b ID      = 1
)

func main() {
	// verbos utilizados (%v, %T, ...) https://pkg.go.dev/fmt
	fmt.Printf("O valor de a é %v\n", a) // 1.2
	fmt.Printf("O tipo de a é %T\n", a)  // float64
	fmt.Printf("O tipo de b é %T\n", b)  // main.ID
}
