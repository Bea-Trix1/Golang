package main

import "fmt"

// Percorrendo um slice com for e range
func main() {
	numeros := []string{"um", "dois", "três"}

	// Range vai percorrer o slice e retornar o índice e o valor
	for c, v := range numeros {
		print(c, " - ", v, "\n")
	}

	// While em Go - Nem existe while em Go, mas podemos simular com for
	for i := 0; i < 10; i++ {
		fmt.Print(i, "\n")
	}

	// Loop infinito - usado para por exemplo colocar um servidor para rodar, consumir uma fila, etc
	for {
		fmt.Print("Loop infinito\n")
		break // Para o loop
	}

}
