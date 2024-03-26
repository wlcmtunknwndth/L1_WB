package main

import "fmt"

func main() {
	num1 := 12314123
	fmt.Println("first num: ", num1)

	num2 := 85648432
	fmt.Println("second num: ", num2)

	//switch
	num1, num2 = num2, num1
	fmt.Println("first num: ", num1)
	fmt.Println("second num: ", num2)

}
