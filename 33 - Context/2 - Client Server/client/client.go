package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // criando um contexto com timeout de 3 segundos
	defer cancel()                                                          // cancelando o contexto após a execução da função

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil) // criando uma requisição com o contexto
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req) // fazendo a requisição
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
