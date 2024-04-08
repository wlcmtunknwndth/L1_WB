package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

func removeElementUnordered(arr []int, ind int) []int {
	if len(arr) == 0 {
		return nil
	}
	if ind > len(arr)-1 {
		return nil
	}

	if ind != len(arr)-1 {
		arr[ind] = arr[len(arr)-1]
	}
	arr = (arr)[:len(arr)-1]
	return arr
}

func main() {
	var slice []int
	gofakeit.Slice(&slice)

	fmt.Println(slice)

	var ind int
	fmt.Scan(&ind)

	slice = removeElementUnordered(slice, ind)
	fmt.Println("after deletion: ", slice)

}
