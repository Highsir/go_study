package main

import "fmt"

func main() {
	s := make([]byte, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println()
	fmt.Println(len(s[2:4]))
	fmt.Println(cap(s[2:4]))
}
