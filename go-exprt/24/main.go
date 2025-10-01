package main

func main() {
	a := 1
	b := 2

	// em go não existe else if
	if a < b {
		println(a)
	} else {
		println(b)
	}

	// operadores de comparação
	// > - maior que
	// < - menor que
	// == - igual
	// >= - maior ou igual
	// <= - menor ou igual
	// && - AND
	// || - OR
	// ! - NOT

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}
}
