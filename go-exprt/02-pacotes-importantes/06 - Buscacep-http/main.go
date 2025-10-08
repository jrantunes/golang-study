package main

import "net/http"

func main() {
	// http - multiplexer - (um recurso com outros recursos(funções) encaixados nele)

	// o handler pode ser escrito assim (não recomendado)
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello world!"))
	// })
	http.HandleFunc("/", BuscaCEPHandler)
	http.ListenAndServe(":8081", nil) // rodando o http server
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound) // envia uma resposta http com um status code (nesse caso 404 not found)
		// http.NotFound(w, r)
		return
	}
	cepParam := r.URL.Query().Get("cep") // recuperando o valor de um query param (?cep=123213213)
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json") // altera o valor de um header
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
}
