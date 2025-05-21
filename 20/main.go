package main

// generics permitem que não seja necessário repetir a mesma função para tipos diferentes ou que não seja necessário o uso de interfaces vazias
// func Soma[T int | float64](m map[string]T) T {
// 	var soma T
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// constraints são tipos especificos que criamos para serem substituidos na tipagem
type Number interface {
	// int | float64 // dessa forma o tipo MyNumber não será aceito por ser um tipo diferente (mesmo sendo um int por debaixo dos panos)
	~int | ~float64 // MyNumber será considerado como int nesse caso
}

type MyNumber int

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// any não permite comparações
// func Compara[T any](a, b T) bool {
// 	if a == b {
// 		return true
// 	}
// 	return false
// }

// comparable aceita apenas valores comparaveis
// compara apenas a igualdade (ele não sabe se uma coisa pode ser maior ou menor que outra)
func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"a": 1000, "b": 2000, "c": 3000}
	m2 := map[string]float64{"a": 100.0, "b": 200.0, "c": 300.0}
	m3 := map[string]MyNumber{"a": 1000, "b": 2000, "c": 3000}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	// println(Compara(10, "10.0")) // erro
	println(Compara(10, 10.0))     // true
	println(Compara(10, 10))       // true
	println(Compara("ana", "ada")) // false
	println(Compara("ana", "ana")) // true
}
