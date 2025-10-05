package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt") // cria um arquivo
	if err != nil {
		panic(err)
	}
	size, err := f.Write([]byte("UAU!, UAU, UAU !!!?? ASDSADASD")) // insere uma bytes no arquivo retornando o novo tamanho dele
	// size, err := f.WriteString("UAU")   // insere uma string no arquivo retornando o novo tamanho dele em bytes
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado! tamanho %d bytes\n", size)
	f.Close() // fecha o arquivo criado

	// leitura
	arquivo, err := os.ReadFile("arquivo.txt") // retorna o arquivo em bytes
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo)) // conteudo do arquivo em string

	// leitura de pouco em pouco (para arquivos maiores que a quantidade de memória disponível)
	arquivo2, err := os.Open("arquivo.txt") // abre o arquivo
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2) // reader para ler o arquivo em buffers
	buffer := make([]byte, 10)          // limitando o buffer em 10 bytes
	for {
		n, err := reader.Read(buffer) // lendo os pedaços do arquivo
		if err != nil {
			break // quando todo o arquivo for lido o loop irá parar
		}
		fmt.Println(string(buffer[:n])) // printando o pedaço do arquivo lido
	}

	err = os.Remove("arquivo.txt") // remove o arquivo
}
