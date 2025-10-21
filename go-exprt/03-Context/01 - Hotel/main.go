package main

import (
	"context"
	"fmt"
	"time"
)

// sistema de reserva de hotel usando contextos
func main() {
	ctx := context.Background() // iniciando um contexto em branco na thread principal
	/*
		WithCancel - podemos cancelar a qualquer momento independente de tempo
		WithDeadline - daqui a quanto tempo podemos cancelar
		WithTimeout - contagem regressiva para o cancelamento
		WithValue - quando passamos um valor
	*/
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) // 3 segundos | Hotel booking cancelled. Timeout reached.
	// ctx, cancel := context.WithTimeout(ctx, 10 * time.Second) // 10 segundos | Hotel booked!
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// select - semelhante a um switch/case | aguarda resultados e quando eles chegam executamos uma ação
	select {
	case <-ctx.Done(): // caso o contexto seja cancelado ou sua regra seja satisfeita
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(5 * time.Second): // caso passe 5 segundos e o context não tenha sido Done()
		fmt.Println("Hotel booked!")
		return
	}
}
