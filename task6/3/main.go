package main

import (
	"fmt"
	"github.com/wlcmtunknwndth/L1_WB/task6/randomChannel"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	chData := randomChannel.FillCh()
	quit := make(chan struct{})
	// finishing goroutine with "for select"
	go func() {
		fmt.Println("Goroutine 3 started")
		for {
			select {
			case data := <-chData:
				fmt.Printf("got data: %s\n", data)
			case <-quit:
				fmt.Println("Goroutine 3 finished")
				return
			}
		}
	}()

	time.Sleep(3 * time.Second)
	quit <- struct{}{}
	time.Sleep(1 * time.Second)

	close(quit)
	close(chData)
	wg.Wait()
}
