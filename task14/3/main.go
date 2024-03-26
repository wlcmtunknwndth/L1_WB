package main

import "fmt"

func determineType(data any) {
	fmt.Printf("the data is of type: %T\n", data)
}

func main() {
	var1 := "idk"
	determineType(var1)
}
