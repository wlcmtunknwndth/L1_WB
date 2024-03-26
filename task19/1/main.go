package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	var str string = "abcdefghijkmnolp"
	var newStr strings.Builder
	newStr.Grow(len(str))
	for i := len(str) - 1; i >= 0; i-- {
		newStr.WriteByte(str[i])
	}
	fmt.Println(newStr.String())
	fmt.Println(time.Since(start))
}
