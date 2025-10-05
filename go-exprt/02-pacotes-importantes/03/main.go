package main

import "fmt"

// func main() {
// 	req, err := http.Get("https://www.google.com")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer req.Body.Close() // defer - esse trecho será executado no final da função
// 	res, err := io.ReadAll(req.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(res))
// }

func main() {
	fmt.Println("primeiro")
	defer fmt.Println("segundo")
	fmt.Println("terceiro")
	// a ordem será primeiro - terceiro - segundo
}
