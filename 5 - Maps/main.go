package main

import "fmt"

func main() {

	salarios := map[string]int{"Beatriz": 1000, "Bea": 100}

	fmt.Println("salarios: ", salarios["Beatriz"])

	sal := make(map[string]int)

	sal["Bea"] = 200

	// underlyne no for, é usado para ignorar o primeiro valor
	for _, s := range sal {
		fmt.Println(s)
	}
}
