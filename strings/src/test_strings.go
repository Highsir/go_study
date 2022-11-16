package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "HasPrefix"
	fmt.Println(strings.HasPrefix(s, "prefix"))
	fmt.Println(strings.HasSuffix(s, "Prefix"))
	fmt.Println(strings.Contains(s, "Prefix"))
	fmt.Println(strings.Index(s, "Prefix"))

	str1 := "aabbccddabcdabcd"
	fmt.Println(strings.Replace(str1, "ab", "a", 1))
	fmt.Println(strings.Count(str1, "ab"))
	fmt.Println(strings.Repeat(s, 3))
	fmt.Println(strings.Fields(str1))
	fmt.Println(strings.Split(str1, "a"))
}
