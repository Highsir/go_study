package main

import (
	"fmt"
	"math"
)

type Point3 struct {
	x, y, z float64
}

func (p Point3) ABS() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func main() {
	p3 := Point3{1, 2, 3}
	fmt.Println(p3.ABS())
}
