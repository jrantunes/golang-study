package matematica

// a função declarada da forma abaixo não será exportada por seu nome ser iniciado com letra minuscula
// quando for iniciado com letra minuscula só será possivel utilizar internamente
// func soma[T int | float64](a, b T) T {
// 	return a + b
// }

// para que metodos/variaveis/strucs/etc.. sejam exportados para fora seu nome deve iniciar com uma letra maiuscula
func Soma[T int | float64](a, b T) T {
	return a + b
}

// var a int = 50 // acessivel somente internamente
// type carro struct {} // acessivel somente internamente
var A int = 50 // acessivel externamente

type Carro struct {
	Marca string // acessivel externamente
	// marca string // acessivel somente internamente
}

// acessivel somente internamente
// func (c Carro) andar() {}

// acessivel externamente
func (c Carro) Andar() {
	println("andou")
}
