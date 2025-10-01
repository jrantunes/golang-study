package main

// com os arquivos go.mod e go.sum conseguimos garantir a versão correta utilizadas em cada pacote
// para instalar um pacote externo basta usar o comando: go get link-do-pacote
// ex: go get golang.org/x/exp/constraints || baixa o pacote de constraints: https://pkg.go.dev/golang.org/x/exp/constraints
// assim que o pacote é instalado ele é adicionado junto a sua versão aos arquivos go.mod e go.sum
// go mod tidy || remove os pacotes que não estão sendo utilizados e baixa os que ainda não foram baixados
// mantendo o go.mod limpo

import (
	"fmt"

	"github.com/google/uuid" // instalado através do go mod tidy || ou via go get github.com/google/uuid
	"github.com/jrantunes/golang-study/matematica"
)

func main() {
	soma := matematica.Soma(10, 30)
	fmt.Println("O resultado é: ", soma)
	fmt.Println(uuid.New())
}
