package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var str string = "abcdefghijkmnolp"
	ans := make([]byte, len(str))

	for i := len(str) - 1; i >= 0; i-- {
		ans[len(str)-i-1] = str[i]
	}
	fmt.Println(string(ans))
	fmt.Println(time.Since(start))

}
