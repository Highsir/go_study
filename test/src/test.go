package main

import "fmt"

const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009
const Log2E = 1 / Ln2 // this is a precise reciprocal
const Billion = 1e9   // float constant
const hardEight = (1 << 100) >> 97

func main() {
	fmt.Println(Ln2)
	fmt.Println(Log2E)
	fmt.Println(Billion)
	fmt.Println(hardEight)
}
