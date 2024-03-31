package main

import "fmt"

func main() {

	salarios := map[string]int{"Beatriz": 1000, "Joao": 100}

	fmt.Println("salarios: ", salarios["Beatriz"])

	sal := make(map[string]int)

	sal["Joao"] = 200

	for _, s := range sal {
		fmt.Println(s)
	}
}
