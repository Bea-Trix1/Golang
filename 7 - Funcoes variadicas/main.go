package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 4, 7, 8, 78, 8, 1000, 54, 45, 4, 100))
}

// Somando numero, ... é utilizado para contar
// a quantidade de parametros que nao sei quantos estão sendo usados
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
