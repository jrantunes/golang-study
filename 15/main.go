package main

func main() {
	// memória -> endereço -> valor
	// o ponteiro é o endereçamento na memória
	// a variável é um local que aponta para um ponteiro que tem um endereço na memória
	// variavel -> ponteiro -> valor
	a := 10
	println(&a)            // endereço na memória de 'a'
	var ponteiro *int = &a // um ponteiro que aponta para o endereço de 'a'
	println(ponteiro)
	println(a)     // 10
	*ponteiro = 20 // altera o valor que está armazenado no endereço de 'a' atraves do ponteiro
	println(a)     // 20
	b := &a        // armazenando o endereço de 'a' em 'b'
	println(b)
	// * - referencing - retorna o valor guardado no endereço da memória
	println(*b) // 20
	*b = 30
	println(*b) // 30
	println(a)  // 30

	// ponteiros são normalmente utilizados porque não queremos que o que estamos criando esteja as vezes em um escopo muito local
	// muitas vezes queremos trabalhar com isso e não importa em qual lugar alguem fez alguma alteração o valor sempre vai estar atualizado
}
