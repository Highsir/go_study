package main

var p string

func main() {
	p = "G"
	print(p)
	f1()
}

func f1() {
	p := "O"
	print(p)
	f2()
}

func f2() {
	print(p)
}
