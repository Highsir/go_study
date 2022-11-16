package main

import (
	"fmt"
	"math"
)

type Square2 struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper2 interface {
	Area() float32
}

func main() {
	var areaIntf Shaper2
	sq1 := new(Square2)
	sq1.side = 5

	if t, ok := areaIntf.(*Square2); ok {
		fmt.Printf("The type of areaIntf is： %T/n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

}

func (sq *Square2) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
