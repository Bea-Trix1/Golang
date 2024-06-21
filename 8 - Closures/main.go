package main

import "fmt"

// Funções anonimas- dentro de uma função eu posso ter outra função. Usada também como handler;
func main() {
	total := func() int {
		return sum(1, 21, 3, 4, 4, 7, 8, 78, 8, 1000, 54, 45, 4, 100) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
