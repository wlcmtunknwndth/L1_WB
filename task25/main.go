package main

import (
	"fmt"
	"time"
)

func sleep(tm time.Duration) {
	//start := time.Now().UnixNano()
	ticker := time.NewTicker(tm)

	<-ticker.C
	ticker.Stop()
}

const second = 1000 * 1000 * 1000

func main() {
	start := time.Now()

	sleep(time.Second)

	fmt.Println(time.Since(start))

}
