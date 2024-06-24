package main

import "fmt"

func CharCount(s string) map[rune]int {
	cont := make(map[rune]int)

	for _, chars := range s {
		cont[chars]++
	}

	return cont
}

func main() {
	s := "Beatriz"

	fmt.Println(CharCount(s))
}
