package main

import "fmt"

func main() {
	var temp = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	var group = make(map[int][]float32)

	for i := range temp {
		group[int(temp[i]/10)*10] = append(group[int(temp[i]/10)*10], temp[i])
	}

	fmt.Printf("map: %+v", group)
}
