package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"sync"
	"time"
)

type Writer struct {
	getX      chan int
	sendXsqrt chan int
}

func main() {
	x := make(chan int)
	xSqrt := make(chan int)

	var wg sync.WaitGroup

	// writer
	wg.Add(3)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Second)
			x <- gofakeit.IntRange(0, 1000)
		}
	}()

	// first conv
	go func() {
		defer wg.Done()
		for {
			select {
			case val := <-x:
				fmt.Printf("received x: %d\n", val)
				xSqrt <- val * val
			}

		}
	}()

	// second conv
	go func() {
		defer wg.Done()
		for {
			select {
			case val := <-xSqrt:
				fmt.Printf("received x^2: %d\n", val)
			}
		}
	}()

	wg.Wait()
}
