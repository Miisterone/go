package main

import "fmt"

func RecursivePower(nb int, power int) int {
	result := 1
	if power < 0 {
		result = 0
	} else if power > 0 {
		result = nb * RecursivePower(nb, power-1)
	}
	return result
}

func main() {
	fmt.Println(RecursivePower(4, 3))
}
