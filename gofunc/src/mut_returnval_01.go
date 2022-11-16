package main

import (
	"fmt"
)

var num1, num2 int = 3, 3

func main() {
	sum := getNum1AndNum2Sum1(num1, num2)
	product := getNum1AndNum2Product1(num1, num2)
	difNum := getNum1AndNum2Dif1(num1, num2)
	printResult1(sum, product, difNum)
}

func printResult1(sum, product, difNum int) {
	fmt.Println(sum, product, difNum)
}

func getNum1AndNum2Dif1(num1 int, num2 int) int {

	return num1 - num2
}

func getNum1AndNum2Product1(num1 int, num2 int) int {

	return num1 * num2
}

func getNum1AndNum2Sum1(num1 int, num2 int) int {

	return num1 + num2
}
