package main

import "fmt"

func main() {
	var mouthNumber int = 6
	switch {
	case mouthNumber < 12 && mouthNumber > 8:
		fmt.Println("mouth is autumn")
	case mouthNumber < 9 && mouthNumber > 5:
		fmt.Println("mouth is summer")
	case mouthNumber < 6 && mouthNumber > 2:
		fmt.Println("mouth is spring")
	}
}
