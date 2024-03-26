package main

import (
	"math/rand"
	"time"
)

const chars = "abcdefghijklmnopqrtstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var justString string

// проблема в том, что в наша ф-ция createHugeString создает переменную не определенного размера, что означает, что может создаться переменная размером меньше 100,
// а наша глобальная переменная в ф-ции создает строку размером 100, поэтому в случае, если будет создана строка размера меньше, чем 100, то программа
// упадет с паникой из-за выхода за пределы слайса v.
func someFunc() {
	v := createHugeString(20)
	justString = v[:100]
}

func createHugeString(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	strBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		strBytes[i] = chars[rnd.Intn(len(chars))]
	}
	return string(strBytes)

}

func main() {
	someFunc()
}
