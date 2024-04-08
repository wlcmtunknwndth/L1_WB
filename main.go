package main

import "fmt"

//Что выведет данная программа и почему?

func change(a *int) {
	t := *a * 2
	a = &t
}

func main() {
	a := 3
	change(&a)
	fmt.Println(a)
}
