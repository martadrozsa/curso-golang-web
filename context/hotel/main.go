package main

import (
	"context"
	"fmt"
	"time"
)

func bookhotel(ctx context.Context)  {
	select {
	case <-ctx.Done():
		fmt.Println("Tempo execido para bookar o quarto")
	case <-time.After(time.Second * 5):
		fmt.Println("Quarto reservado com sucesso")
	}
}

func main()  {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()
	bookhotel(ctx)
}
