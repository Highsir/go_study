package main

import "fmt"

func main() {
	var c = Adder2()
	fmt.Print(c(1), " - ")
	fmt.Print(c(20), " - ")
	fmt.Print(c(300))
}
func Adder2() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
