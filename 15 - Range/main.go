package main

import "fmt"

// O range do laço for itera sobre uma slice ou map.
// Ao variar sobre uma slice, dois valores são retornados para cada iteração. O primeiro é o índice, o segundo uma cópia do elemento daquele índice.

var pow = []int{1, 2, 4, 8, 16, 32, 62, 128} // armazena a potência de 2

func main() { // range itera sobre a slice pow - ou seja, ele vai percorrer o pow
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
