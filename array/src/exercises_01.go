package main

import "fmt"

func main() {
	str := "aabbcc"
	new_str := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(new_str)
}
