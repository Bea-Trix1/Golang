package main

// Descrição do desafio:

//Exercício: Slices
// Implemente Pic. Ela deve retornar uma slice de comprimento dy, cada elemento do qual é uma slice de dx 8-bit inteiros sem sinal.
// Quando você executa o programa, ele irá exibir a sua imagem, interpretando os números inteiros como escala dos valores de cinza.
// A escolha da imagem é com você. Funções interessantes incluem x^y, (x+y)/2, e x*y.
// (Você precisa usar um loop para alocar cada []uint8 dentro do [][]uint8.)
// (Utilize uint8(intValue) para converter entre os tipos.)

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)

	// Itera sobre cada linha da imagem
	for y := 0; y < dy; y++ {
		// Cria uma slice para a linha de comprimento dx
		img[y] = make([]uint8, dx)

		// Itera sobre cada coluna da linha
		for x := 0; x < dx; x++ {
			// Calcula o valor do pixel usando (x + y) / 2
			img[y][x] = uint8((x + y) / 2)
		}
	}

	return img
}

func main() {
	pic.Show(Pic)
}
