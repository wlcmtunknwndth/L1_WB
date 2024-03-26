package main

import (
	"fmt"
	"reflect"
)

func determineType(data interface{}) {
	fmt.Println("the type is: ", reflect.TypeOf(data))
}

func main() {
	var1 := make(chan []interface{})

	determineType(var1)
}
