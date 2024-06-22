package main

import "fmt"

func main() {
	v := Estrutura{1, "Algo"}
	v.Valor = 10
	fmt.Println(v.Valor)
	v.Produto = "Sabonete" // Campos Struct são acessados ​​através de um ponto.
	fmt.Println(v.Produto)
}

// A struct é uma coleção de campos.
type Estrutura struct {
	Valor   int
	Produto string
}
