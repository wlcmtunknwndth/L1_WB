package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

func Intersect(set1, set2 map[int]struct{}) map[int]struct{} {
	var ans = make(map[int]struct{})

	for key := range set1 {
		if _, ok := set2[key]; ok {
			ans[key] = struct{}{}
		}
	}

	return ans
}

func randomSet() map[int]struct{} {
	rndLen := int(gofakeit.Uint8() % 17)
	var ans = make(map[int]struct{}, rndLen)

	for i := 0; i < rndLen; i++ {
		ans[int(gofakeit.Int8()%19)] = struct{}{}
	}
	return ans
}

func main() {
	//set1 := map[int]struct{}{
	//	1: {},
	//	2: {},
	//	3: {},
	//	4: {},
	//}
	//set2 := map[int]struct{}{
	//	1: {},
	//	3: {},
	//	4: {},
	//	5: {},
	//}
	set1 := randomSet()
	fmt.Println("set #1: ", set1)

	set2 := randomSet()
	fmt.Println("set #2: ", set2)

	fmt.Printf("%+v", Intersect(set1, set2))
}
