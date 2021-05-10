package main

import "fmt"

func Sqrt(nb int) int {

	if nb == 1 {
		return 1
	}

	for i := 0; i <= nb; i++ {
		if i*i == nb {
			return i
		}

	}
	return 0
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
