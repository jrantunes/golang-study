package main

// composição de structs
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // a struct Endereço está compondo a struct Cliente
	// Address Endereco  // aqui temos uma propriedade de tipo Endereço
}

func main() {
	cliente := Cliente{
		Nome:  "Carlos",
		Idade: 20,
		Ativo: true,
	}

	cliente.Cidade = "São Paulo"
	// ou
	// cliente.Endereco.Cidade = "São Paulo"
	println("Hello 12")
}
