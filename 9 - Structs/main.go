package main

import "fmt"

func main() {
	fmt.Println(Estrutura{1, "bia"})
}

// A struct é uma coleção de campos.
type Estrutura struct {
	Valor   int
	Produto string
}
