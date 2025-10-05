package main

// escopo global // não precisam ser utilizadas após a declaração
const msg = "Hello world" // constante
var b bool                // variavel boolean (inferido por padrão como false)

// definição de multiplas variaveis
// var (
// 	a bool
// 	c int     // inferido por padrão como 0
// 	d string  // inferido por padrão como ''
// 	e float64 // inferido por padrão como +0.000000e+000
// )

// declaração e atribuição
var (
	a bool    = true
	c int     = 10
	d string  = "test"
	e float64 = 1.2
)

func main() {
	// escopo local
	// var escopada string // precisam ser utilizadas dentro do escopo após a declaração

	// b = 10 // erro de tipagem (b foi definida como boolean)

	// declaração e atribuição com inferencia de tipo
	a := "Teste" // o tipo string já é inferido
	// := é feito somente na primeira vez em que a variavel é criada
	println(a)
	a = "Teste 2"
	println(a)
}

// func x() {
// println(escopada) // não acessivel
// }
