package main

import "fmt"

func main() {
	fmt.Println(b())
}

func b() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
