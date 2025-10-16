package main

import "net/http"

// FileServer serve para servir arquivos estáticos (css, imagens, html, etc...)
func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello blog :)"))
	})
	http.ListenAndServe(":8080", mux)
}
