package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

func removeElementOrdered(slice *[]int, ind int) {
	*slice = append((*slice)[:ind], (*slice)[ind+1:]...)
}

func main() {
	var slice []int
	gofakeit.Slice(&slice)

	fmt.Println(slice)

	var ind int
	fmt.Println("choose index to delete from slice: ")
	fmt.Scan(&ind)

	slice = append(slice[:ind], slice[ind+1:]...)

	fmt.Println("After deletion: ", slice)
}
