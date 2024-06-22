package main

import (
	"fmt"
	"math"
)

// Dado um número x, queremos encontrar o número z para o qual z² é quase o x.

// Sqrt calcula a raiz quadrada de x usando o método iterativo
func Sqrt(x float64) float64 {
	// Estimativa inicial
	z := 1.0

	// Iterar até a aproximação ser precisa ou atingir o número máximo de iterações
	for i := 0; i < 10; i++ {
		valor := z - (z*z-x)/(2*z)
		fmt.Printf("Iteração %d: z = %f\n", i+1, z)
		z = valor
	}
	return z
}

func main() {
	// Testar a função Sqrt com vários valores de x
	for i := 1; i <= 10; i++ {
		x := float64(i)
		approx := Sqrt(x)
		actual := math.Sqrt(x)
		fmt.Printf("x = %d, Sqrt(x) = %f, math.Sqrt(x) = %f, diff = %f\n", i, approx, actual, math.Abs(approx-actual))
	}
}
