package main

import "fmt"

//Что выведет данная программа и почему?

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice)
	}(slice)
	fmt.Println(slice)
}
