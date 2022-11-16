package main

func main() {
	F1("a", "b", "c")
}

func F1(s ...string) {
	F2(s...)
	F3(s)
}

func F2(s ...string) {
	for _, str := range s {
		println(str)
	}
}
func F3(s []string) {
	println(s)
}
