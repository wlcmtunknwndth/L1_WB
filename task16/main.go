package main

import "fmt"

func partition(arr []int, l, r int) int {
	//src := rand.NewSource(time.Now().UnixNano())
	//rnd := rand.New(src)
	pivot := arr[(l+r)/2]

	for {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l >= r {
			return r
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
}

func quickSort(arr []int, l, r int) {
	if l < r {
		//pivot := (l + r) / 2
		p := partition(arr, l, r)
		quickSort(arr, l, p)
		quickSort(arr, p+1, r)
	}
}

func main() {
	testArr := []int{1, 2342, 12, 2357, 1213, 31, 99, 214, 2, 67, 84}
	quickSort(testArr, 0, len(testArr)-1)
	fmt.Println(testArr)
}
