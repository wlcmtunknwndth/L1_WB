package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// just waiting for program to finish itself
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
		for {
			if num := rand.Intn(300) % 10; num == 0 {
				fmt.Printf("reached stop condition: %d\n", num)
				break
			}
			fmt.Println("g")
		}
		fmt.Println("Goroutine 1 finished")
	}()
	wg.Wait()
}
