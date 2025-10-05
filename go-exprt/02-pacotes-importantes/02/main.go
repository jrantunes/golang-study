package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com") // stream com a resposta da requisição
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body) // lendo a resposta
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	req.Body.Close() // fechando a stream
}
