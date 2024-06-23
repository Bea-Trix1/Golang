package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow { // _ serve para ignorar o indice(posição) da slice. O value representa o valor atual da iteração range.
		fmt.Printf("%d\n", value)
	}
}
