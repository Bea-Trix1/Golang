package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", 1234) // contexto.WithValue permite passarmos um contexto com chave e valor
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	userID := ctx.Value("userID") // recuperando o valor da chave userID
	fmt.Println("userID: ", userID)
}
