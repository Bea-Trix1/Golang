package main

import "fmt"

//Uma slice, é dinamicamente redimensionada
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	fmt.Println("---------------------------------------")

	// Uma slice não armazena todos os dados, apenas descreve uma seção de uma matriz subjacente.
	// Alterando os elementos de uma slice modifica os elementos correspondentes da sua matriz subjacente.
	nomes := [4]string{
		"Morango",
		"Manga",
		"Banana",
		"Kiwi",
	}

	fmt.Println(nomes)
	a := nomes[0:2]
	b := nomes[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println(nomes)
}
