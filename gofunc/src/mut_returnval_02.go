package main

import (
	"fmt"
)

var num3, num4 int = 3, 3

func main() {
	sum := getNum1AndNum2Sum(num1, num2)
	product := getNum1AndNum2Product(num1, num2)
	difNum := getNum1AndNum2Dif(num1, num2)
	printResult(sum, product, difNum)
}

func printResult(sum, product, difNum int) {
	fmt.Println(sum, product, difNum)
}

func getNum1AndNum2Dif(num1 int, num2 int) int {
	difNum := num1 - num2
	return difNum
}

func getNum1AndNum2Product(num1 int, num2 int) int {
	product := num1 * num2
	return product
}

func getNum1AndNum2Sum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}
