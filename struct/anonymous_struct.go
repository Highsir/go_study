package main

import "fmt"

type M struct {
	x int
	y float32
}
type N struct {
	x int
	y float32
}

type C struct {
	M
	N
}

func main() {
	c := C{M{1, 2.0}, N{3, 4.0}}
	fmt.Println(c)
}
