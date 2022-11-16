package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d"}
	fmt.Println(a)
	for i, j := range a {
		fmt.Println("Array item", i, j, "is", a[i])
	}
}
