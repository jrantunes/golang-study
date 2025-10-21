package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	// configurando uma request
	// client http !== objeto de request
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	// adicionando mais configurações na request
	req.Header.Set("Accept", "application/json") // aceito application/json no retorno da request
	// juntando a request com o client http
	resp, err := c.Do(req)
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
