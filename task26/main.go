package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const chars = "abcdefghijklmnopqrtstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	strBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		strBytes[i] = chars[rnd.Intn(len(chars))]
	}
	return string(strBytes)

}

func CharsUnique(str string) bool {
	if len(str) <= 1 {
		return true
	}
	str = strings.ToLower(str)

	checker := make(map[int32]struct{})

	for _, ch := range str {
		if _, ok := checker[ch]; ok {
			return false
		}
		checker[ch] = struct{}{}
	}

	return true
}

func main() {
	str := RandomString(10)
	fmt.Printf("the string '%s' is unique? %t", str, CharsUnique(str))
}
