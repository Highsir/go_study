package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	f1 bool   "An important answer"
	f2 string "The name of the thing"
	f3 int    "How much there are"
}

func main() {
	tt := TagType{true, "hello", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {

	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
