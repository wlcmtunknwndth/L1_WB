package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 121, 14, 1231231, 435421, 1213, 531, 99, 51}
	fmt.Println("using WaitGroup: ")
	WaitGroup(arr)
	fmt.Println("\nusing Mutex")
	Mutex(arr)
}

func WaitGroup(arr []int) {
	var wg sync.WaitGroup

	wg.Add(len(arr))

	start := time.Now()

	for _, val := range arr {
		go func(val int) {
			fmt.Printf("%d ", val*val)
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

func Mutex(arr []int) {
	var mtx sync.Mutex

	start := time.Now()

	for _, val := range arr {
		go func(val int) {
			fmt.Printf("%d ", val*val)
			mtx.Unlock()
		}(val)
		mtx.Lock()
	}

	fmt.Println(time.Since(start))
}
