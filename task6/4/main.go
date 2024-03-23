package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

func main() {
	//with context

	quit := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		fmt.Println("goroutine started")
		for {
			select {
			case <-ctx.Done():
				quit <- struct{}{}
				fmt.Println("goroutine finished")
				return
			default:
				fmt.Printf("Proceeded: %s\n", gofakeit.Emoji())
			}
		}
	}(ctx)

	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	<-quit
}
