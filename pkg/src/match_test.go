package main

import (
	"fmt"
	"regexp"
)

func main() {
	ok, _ := regexp.Match("aabb", []byte("ab"))
	fmt.Println(ok)
}
