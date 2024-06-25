package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1 // Inicializando os dois primeiros números da sequência

	return func() int {
		result := a   // O próximo número da sequência é armazenado em 'result'
		a, b = b, a+b // Atualizando os valores de 'a' e 'b' para a próxima iteração
		return result // Retornando o próximo número da sequência de Fibonacci
	}
}

func main() {
	f := fibonacci()

	// Imprimindo os primeiros 10 números da sequência de Fibonacci
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
