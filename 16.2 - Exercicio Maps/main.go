package main

import (
	"fmt"
	"strings"
)

// Exercício: Maps
// Implementar WordCount. Ele deve retornar um map das contagens de cada "word" na string s. A função wc.Test executa um conjunto de testes contra a função fornecida e imprime o sucesso ou falha.
// Você pode achar strings.Fields útil.

func WordCount(s string) map[string]int {
	// Dividindo a string em palavras
	palavras := strings.Fields(s) // strings.Fields(s): Esta função divide a string s em palavras, retornando um slice de strings.

	// Criando um map para armazenar as contagens
	contador := make(map[string]int)

	// Iterando sobre as palavras e atualizando as contagens
	for _, palavra := range palavras { //Itera sobre cada palavra no slice palavra.
		contador[palavra]++ // Incrementa a contagem da palavra no map.
	}
	//Retornando o map
	return contador
}

func main() {
	s := "Eu amo a Laura barros!"
	fmt.Println(WordCount(s))
}
