package main

import "fmt"

func main() {
	s1_form := []int{1, 2, 3}
	s1_to := make([]int, 10)
	
	n := copy(s1_to, s1_form)
	fmt.Println(s1_to)
	fmt.Printf("Copied %d elements\n", n)

	s1_3 := []int{1, 2, 3}
	s1_3 = append(s1_3, 4, 5, 6)
	fmt.Println(s1_3)
}
