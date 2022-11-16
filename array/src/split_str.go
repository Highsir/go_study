package main

import "fmt"

func split_str(str string, i int) (string, string) {
	var m []byte = []byte(str)
	var s1 []byte = m[0:i]
	var s2 []byte = m[i:]
	return string(s1), string(s2)
}

func main() {
	var str = "abcde"
	num := 2
	str1, str2 := split_str(str, num)
	fmt.Println(str1, str2)
}
