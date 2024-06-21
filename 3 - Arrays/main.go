package main

import "fmt"

type id int

var (
	f id = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	//contando quantas posiçoes tem meu array
	fmt.Println(len(meuArray))

	//percorrendo meu array - range é utilizado para percorrer.
	for i, v := range meuArray {
		fmt.Printf("o valor do indice %d é: %d \n", i, v)
	}
}
