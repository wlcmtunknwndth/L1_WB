package main

import (
	"fmt"
	"slices"
)

func cmp(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func main() {
	var testArr []int
	//gofakeit.Slice(&testArr)
	testArr = []int{1, 231, 4321, 3426, 654, 3, 51211, 231, 213214, 214, 7695, 5721, 75632, 125, 54311, 12}
	slices.SortFunc(testArr, cmp)
	fmt.Println(testArr)

	fmt.Printf("number %d is at the index %d", 51211, BinarySearch(&testArr, 51211))
}

func BinarySearch(arr *[]int, key int) int {
	l := 0
	r := len(*arr) - 1

	for l <= r {
		mid := (l + r) >> 1
		midVal := (*arr)[mid]

		if midVal < key {
			l = mid + 1
		} else if midVal > key {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
