package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"sync"
	"time"
)

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) GetValue() int {
	return c.value
}

func main() {
	var cntr Counter

	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			emji := gofakeit.Emoji()
			select {
			case ch <- emji:
				if cntr.GetValue() > 5 {
					return
				}
				time.Sleep(time.Second)
				fmt.Printf("sent data: %s\n", emji)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case data := <-ch:
				if cntr.GetValue() > 5 {
					return
				}
				fmt.Printf("received data: %s\n", data)
				cntr.Increment()
			}
		}
	}()

	wg.Wait()
	close(ch)
	fmt.Printf("Packets received: %d\n", cntr.GetValue())
}
