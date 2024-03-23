package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"sync"
	"time"
)

func main() {
	// closing channel
	chData := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ch chan string) {
		defer wg.Done()
		fmt.Println("Goroutine started")
		for {
			select {
			case data, ok := <-chData:
				if !ok {
					fmt.Println("Goroutine finished")
					return
				}
				fmt.Printf("received message: %s\n", data)
			}
		}
	}(chData)

	for i := 0; i < 3; i++ {
		chData <- gofakeit.Emoji()
		time.Sleep(time.Second)
	}

	close(chData)
	wg.Wait()
}
