package main

// sempre que iniciamos um projeto em go devemos criar um modulo para poder trabalhar com pacotes usando o comando go mod init nome-do-modulo
// por convenção o nome-do-modulo geralmente é usado o link do repositório/pacote utilizado na hora da importação
// ex: go mod init github.com/jrantunes/golang-study  ||  para importar algum pacote seria: github.com/jrantunes/golang-study/nome-do-pacote
// para instalar/atualizar dependencias devemos usar o comando go mod tidy

import (
	"fmt"

	"github.com/jrantunes/golang-study/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Println("Resultado: ", soma) // Resultado:  30
}
