package main

import "net/http"

func main() {
	// http - multiplexer - (um recurso com outros recursos(funções) encaixados nele)

	// o handler pode ser escrito assim (não recomendado)
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello world!"))
	// })
	http.HandleFunc("/", BuscaCEP)
	http.ListenAndServe(":8081", nil) // rodando o http server
}

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
