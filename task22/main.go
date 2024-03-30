package main

import (
	"fmt"
	"math/big"
)

func sum(a, b string) string {
	op1, _ := new(big.Float).SetString(a)
	op2, _ := new(big.Float).SetString(b)
	res := new(big.Float).Add(op1, op2)
	//fmt.Println("The sum is ", res)
	return res.String()
}

func sub(a, b string) string {
	op1, _ := new(big.Float).SetString(a)
	op2, _ := new(big.Float).SetString(b)
	res := new(big.Float).Sub(op1, op2)
	return res.String()
}

func multiply(a, b string) string {
	op1, _ := new(big.Float).SetString(a)
	op2, _ := new(big.Float).SetString(b)

	res := new(big.Float).Mul(op1, op2)
	return res.String()
}

func divide(a, b string) string {
	op1, _ := new(big.Float).SetString(a)
	op2, _ := new(big.Float).SetString(b)
	res := new(big.Float).Quo(op1, op2)
	return res.String()
}

func main() {
	var var1, var2 string
	_, err := fmt.Scan(&var1, &var2)
	if err != nil {
		fmt.Printf("Error happened: %v", err)
		return
	}
	fmt.Println("addition result: ", sum(var1, var2))
	fmt.Println("subtraction result: ", sub(var1, var2))
	fmt.Println("multiply result: ", multiply(var1, var2))
	fmt.Println("division result: ", divide(var1, var2))

}
