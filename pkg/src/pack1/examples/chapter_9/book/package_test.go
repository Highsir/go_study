package main

import (
	"awesomeProject/pkg/src/pack1"
	"fmt"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Println(test1)
	fmt.Println(pack1.Pack1Int)
}
