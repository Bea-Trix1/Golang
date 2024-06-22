package main

import (
	"fmt"
	"strings"
)

func main() {
	// Criando as colunas
	tabuleiro := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// Os jogadores escolhem as posições.
	tabuleiro[0][0] = "X"
	tabuleiro[2][2] = "O"
	tabuleiro[1][2] = "X"
	tabuleiro[1][0] = "O"
	tabuleiro[0][2] = "X"

	for i := 0; i < len(tabuleiro); i++ {
		fmt.Printf("%s\n", strings.Join(tabuleiro[i], " "))
	}
}
