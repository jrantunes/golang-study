package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

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

// retorna o ponteiro para a variavel original e não uma cópia
func BuscaCEP(cep string) (*ViaCEP, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data ViaCEP
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
