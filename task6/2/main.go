package main

import (
	"fmt"
	"github.com/wlcmtunknwndth/L1_WB/task6/randomChannel"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// finishing with "foreach"
	chData := randomChannel.FillCh()
	go func(ch chan string) {
		defer wg.Done()
		fmt.Println("Goroutine 3 started")
		for val := range ch {
			fmt.Printf("getting data: %s\n", val)
		}
		fmt.Println("Goroutine 3 finished")
	}(chData)

	close(chData)
	wg.Wait()
}
