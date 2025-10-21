package main

import (
	"io"
	"net/http"
	"time"
)

// tempo limite para a execução da requisição
// estorar erro caso esse limite seja excedido
func main() {
	c := http.Client{Timeout: time.Second} // objeto http com timeout de 1 segundo (se esse timeout for excedido irá ocorrer um context deadline exceeded)
	// c := http.Client{Timeout: time.Microsecond} // exemplo com erro
	resp, err := c.Get("http://google.com")
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
