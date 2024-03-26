package main

import "fmt"

func determineType(data any) {
	switch v := data.(type) {
	case int:
		fmt.Println("Interface is an integer")
	case string:
		fmt.Println("Interface is a strings")
	case bool:
		fmt.Println("Interface is a boolean")
	case chan interface{}:
		fmt.Println("Interface is a channel interface")
	case chan string:
		fmt.Println("Interface is a chan interface")
	default:
		fmt.Println("Interface is not determined: %T\n", v)
	}
}

func main() {
	var1 := make(chan interface{})
	determineType(var1)

	var2 := make(chan string)
	determineType(var2)

	var3 := make(chan []int)
	determineType(var3)
}
