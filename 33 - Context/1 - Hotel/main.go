package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3) // definindo timeout do contexto de 3 segundos
	defer cancel()                                         // fechando o contexto quando a função terminar
	BookHotel(ctx)                                         // chamando a função que simula a reserva de hotel passando o contexto
}

func BookHotel(ctx context.Context) {
	select { // select é um switch que só funciona com canais, canais sao como filas de mensagens
	case <-ctx.Done(): // verifica se o contexto foi cancelado, entra nesse case apos 3 segundos
		fmt.Println("Reserva de hotel cancelada. Tempo excedido.")
		return
	case <-time.After(5 * time.Second): // simula a reserva de hotel, esperando 5 segundos
		fmt.Println("Hotel reservado!")
	}
}
