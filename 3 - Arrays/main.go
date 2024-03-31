package main

import "fmt"

type id int

var (
	f id = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	fmt.Println(len(meuArray))

	for i, v := range meuArray {
		fmt.Printf("o valor do indice %d é: %d \n", i, v)
	}
}
