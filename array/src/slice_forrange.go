package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 4)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4
	fmt.Println(len(slice1[1:1]))
	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is %d\n", ix, value)
	}
	for _, value := range slice1 {
		value *= 2
	}
	fmt.Println(slice1)
}
