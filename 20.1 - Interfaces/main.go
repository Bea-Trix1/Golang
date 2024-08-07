package main

import "fmt"

// Um tipo implementa uma interface através da implementação dos métodos.
//  Não há declaração explícita de intenções, não há palavra-chave "implements".

// Interfaces implícitas dissociam-se da definição de uma interface de sua implementação,
//  que pode então aparecer em qualquer pacote sem pré-arranjamento.

type I interface {
	M()
}

type T struct {
	S string
}

//Esse método significa que o type T implementa a interface I,
// Mas não precisamos declarar explicitamente isso.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"olá"}
	i.M()
}
