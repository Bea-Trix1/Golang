package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Quando é sabádo?")
	hoje := time.Now().Weekday()
	switch time.Saturday {
	case hoje + 0:
		fmt.Println("Hoje.")
	case hoje + 1:
		fmt.Println("Amanhã.")
	case hoje + 2:
		fmt.Println("Em dois dias.")
	default:
		fmt.Println("Muito distânte.")
	}
}
