package main

import "net/http"

// http.ListenAndServe - sobe um multiplexer(mux) padrão quando passamos nil (http.ListenAndServe(":8080", nil)) - (um componente que attachamos rotas nele)
// para ter mais de um server rodando em uma porta diferente, para evitar que um pacote externo adicione um handler no nosso servidor e para termos
// mais controle, garantir que os handlers attachados sejam os que nós mesmos attachamos, podemos criar nosso proprio ServeMux(multiplexer)

func main() {
	mux := http.NewServeMux() // responsável por receber nossos handlers
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // attachando um handler no mux
	// 	w.Write([]byte("Hello World!"))
	// })
	mux.HandleFunc("/", HomeHandler)
	// handler implementa uma interface que torna obrigatório o uso de um metodo ServeHTTP
	//ex:
	mux.Handle("/blog", blog{}) // uma struct para cada endpoint | essa forma é mais customizavel
	// ex customização:
	mux.Handle("/blog2", blogWithTitle{title: "Hello blog2"})
	http.ListenAndServe(":8081", mux) // subindo um webserver utilizando o mux criado acima
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello home"))
}

type blog struct{}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello blog"))
}

type blogWithTitle struct {
	title string
}

func (b blogWithTitle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
