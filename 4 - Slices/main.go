package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}

	fmt.Printf("len %d cap %d %v\n", len(s), cap(s), s)

	fmt.Printf("len %d cap %d %v\n", len(s[:0]), cap(s), s[:0])

	fmt.Printf("len %d cap %d %v\n", len(s[:4]), cap(s), s[:4])

	fmt.Printf("len %d cap %d %v\n", len(s[2:]), cap(s), s[2:])

	s = append(s, 12)
	fmt.Printf("len %d cap %d %v\n", len(s[2:]), cap(s[2:]), s[2:])
}
