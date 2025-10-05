package main

// aloca uma cópia da variável recebida na memória
func soma(a, b int) int {
	a = 50 // apenas o valor da cópia da variável é alterado e a original permanece como estava antes
	return a + b
}

// recebe e manipula os valores originais através do endereço deles
func somaComPonteiro(a, b *int) int {
	// o ponteiro é necessário somente se quisermos trabalhar com valores originais mutáveis
	*a = 50
	return *a + *b // retorna a soma dos valores armazenados na memória
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	println(soma(minhaVar1, minhaVar2))              // manda cópias das variaveis para a função
	println(minhaVar1)                               // o valor de 'minhaVar1' permanece como 10
	println(somaComPonteiro(&minhaVar1, &minhaVar2)) // manda os endereços das variáveis
	println(minhaVar1)                               // o valor é alterado para 50
}
