package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(summarize(&arr), time.Since(start))
}

func summarize(arr *[]int) int {
	var ans int
	var wg sync.WaitGroup
	wg.Add(len(*arr))
	for _, val := range *arr {
		go func(ans *int, val int) {
			*ans += val * val
			wg.Done()
		}(&ans, val)
	}
	wg.Wait()
	return ans
}
