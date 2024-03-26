package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

type Set struct {
	data map[string]struct{}
}

func New(arr *[]string) *Set {
	data := make(map[string]struct{}, len(*arr))

	for i := range *arr {
		data[(*arr)[i]] = struct{}{}
	}

	return &Set{
		data: data,
	}
}

func randomStrings() []string {
	rndLen := int(gofakeit.Uint8()) % 17
	var ans = make([]string, rndLen)

	for i := 0; i < rndLen; i++ {
		ans[i] = gofakeit.Emoji()
	}
	return ans
}

func main() {
	// generate random []string
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("as slice: ", strings)

	set := New(&strings)
	fmt.Printf("as set: %+v\n", *set)

	strings = randomStrings()
	fmt.Println("as slice: ", strings)

	set = New(&strings)
	fmt.Printf("as set: %+v\n", *set)
}
