package main

import "fmt"

func IncrementBit(num int, pos uint16) int {
	return num | (1 << (pos - 1))
}

func DecrementBit(num int, pos uint16) int {
	return num & (^(1 << (pos - 1)))
}

func main() {
	num := 126
	fmt.Printf("Was: %b\n", num)

	fmt.Printf("IncrementBit: %b\n", IncrementBit(num, 1))

	fmt.Printf("DecrementBit: %b\n", DecrementBit(num, 4))
}
