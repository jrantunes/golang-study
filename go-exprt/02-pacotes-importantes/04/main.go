package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	// tags são anotações com chave:valor utilizadas nas structs em que bibliotecas podem tomar decisões baseadas nelas
	Numero int `json:"n"` // no Marshal Numero será n e no json serializado na hora do Unmarshal a chave deve estar como n para que o bind seja feito em Numero
	Saldo  int `json:"s"`
	// json:"-" - ignora a prop
}

func main() {
	conta := Conta{
		Numero: 1,
		Saldo:  1212,
	}
	res, err := json.Marshal(conta) // Marshal - serializa a struct em um json e retorna ele em bytes
	if err != nil {
		println(err)
	}
	println(string(res)) // {"n":1,"s":1212}

	err = json.NewEncoder(os.Stdout).Encode(conta) // Encoder - recebe um valor, faz o encoding (processo de serialização) e retorna para algum outro lugar (arquivo, web server, terminal (stdout))
	if err != nil {
		println(err)
	}
	// já exibe no terminal o json

	// jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	// var contaX Conta
	// err = json.Unmarshal(jsonPuro, &contaX) // Unmarshal - pega um json serializado e o transforma em uma struct atribuindo seus valores em uma variável
	// if err != nil {
	// 	println(err)
	// }
	// println(contaX.Saldo) // 200

	// caso o json não satisfaça a struct
	jsonPuro2 := []byte(`{"num":2,"sal":200}`)
	var contaX2 Conta
	err = json.Unmarshal(jsonPuro2, &contaX2)
	if err != nil {
		println(err)
	}
	println(contaX2.Saldo) // 0

	// caso o json satisfaça as tags
	jsonPuro3 := []byte(`{"n":2,"s":200}`)
	var contaX3 Conta
	err = json.Unmarshal(jsonPuro3, &contaX3)
	if err != nil {
		println(err)
	}
	println(contaX3.Saldo) // 200
}
