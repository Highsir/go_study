package main

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// OK
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	// not ok
	//z := make(Bar)
	//(*z).thingOne = "hello"
	//(*z).thingTwo = 1

	// ok
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	//not ok
	u := new(Foo)
	(*u)["x"] = "goodbye"
	(*u)["y"] = "world"
}
