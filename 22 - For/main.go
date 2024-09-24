package main

// Percorrendo um slice com for e range
func main() {
	numeros := []string{"um", "dois", "três"}

	// Range vai percorrer o slice e retornar o índice e o valor
	for c, v := range numeros {
		print(c, " - ", v, "\n")
	}

}
