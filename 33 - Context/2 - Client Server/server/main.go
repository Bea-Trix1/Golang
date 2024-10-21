package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)     // chamando a função handler
	http.ListenAndServe(":8080", nil) // iniciando o servidor
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // pegando o context que veio da requisição
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second): // aguardando 5 segundos para responder
		log.Println("Request processada com sucesso.")
		w.Write([]byte("Request processada com sucesso."))

	case <-ctx.Done(): // caso o contexto seja cancelado
		log.Println("Request cancelada pelo cliente.")
	}
}
