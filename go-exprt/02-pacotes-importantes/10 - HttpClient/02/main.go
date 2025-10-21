package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "claudio"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar) // URL, Content-Type, dados a serem lidos por um Reader (um Buffer nesse caso)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil) // envia uma cópia de resp.Body para o os.Stdout (writer)
	// io.CopyBuffer(os.Stdout, resp.Body, make([]byte, 1024)) // CopyBuffer com tamanho definido
}
