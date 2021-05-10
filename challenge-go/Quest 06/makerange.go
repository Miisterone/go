package main

import (
	"fmt"
)

func MakeRange(min, max int) []int {

	if max <= min {
		return []int(nil)
	}
	s := make([]int, max-min)

	for i := range s {
		s[i] = min + i
	}
	return s

}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
