package main

func main() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	println(s2)
}
