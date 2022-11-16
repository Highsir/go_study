package main

import "fmt"

func main() {
	var str string = Bytes("Google")
	fmt.Println(str)
}

func Bytes(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-i-1]
	}
	return string(r)
}
