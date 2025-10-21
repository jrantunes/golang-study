package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

// contextos - permitem que a gente passe as informações deles para diversas chamadas dentro do sistema
// também temos a opção de fazer com que esses contextos sejam cancelados
// quando um contexto é cancelado, paramos a operação na hora para evitar gasto de tempo

func main() {
	ctx := context.Background() // contexto vazio
	// criaremos um contexto com um timeout de 1 segundo
	// qualquer coisa que estiver utilizando esse contexto irá parar de executar caso esse timeout exceda
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	// ctx, cancel := context.WithCancel(ctx) // WithCancel - não possui nenhuma regra. ele é parado apenas quando a função cancel for executada
	// existe duas formas de dizer quando esse contexto foi cancelado (quando o timeout exceder e quando a função cancel for chamada)
	defer cancel() // esse contexto sempre de alguma forma ele será cancelado. seja pelo timeout excedido ou no final da execução do programa
	// request com contexto
	// a request abaixo só vai continuar executando de acordo com a regra passada no contexto ou na execução do cancel
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req) // quando não quisermos criar um httpClient do zero, podemos usar o httpClient padrão
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
