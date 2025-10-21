package main

import (
	"log"
	"net/http"
	"time"
)

// teremos 5 segundos para processar a nossa requisição
// retornaremos para o usuário que a requisição foi processada ou cancelada

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // context da requisição
	log.Println("Requisição iniciada!")
	defer log.Println("Requisição finalizada!")
	select {
	case <-time.After(5 * time.Second): // quando passar 5 segundos
		// imprime no command line stdout
		log.Println("Requisição processada com sucesso!")
		// imprime no client
		w.Write([]byte("Requisição processada com sucesso!"))
	case <-ctx.Done(): // quando a requisição for cancelada pelo client
		log.Println("Requisição cancelada pelo cliente!")
	}
}
