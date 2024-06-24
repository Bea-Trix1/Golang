package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex // Variavel m recebe um map com tipo da chave em string e o valor da chave Vertex da struct acima

func main() {
	m = make(map[string]Vertex) // fazendo o m receber uma slice
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967, // atribuindo valores no map e na slice
	}
	fmt.Println(m["Bell Labs"])
}
