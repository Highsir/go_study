package main

import "fmt"

type employee struct {
	salary float32
}

func main() {
	em := employee{100}
	fmt.Println(em.salary * 1.2)
}
