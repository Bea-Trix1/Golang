package main

import "fmt"

func main() {
	i, j := 12, 2701

	p := &i //Ponteiro para o valor de i
	fmt.Println(*p)

	*p = 21 // Setando o valor de i para 21 atráves do ponteiro
	fmt.Println(i)

	p = &j         // Ponteiro para o valor de j
	*p = *p / 37   // dividindo o valor de j pelo valor do ponteiro de p
	fmt.Println(j) // novo valor do j

}
