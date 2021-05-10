package main

import "fmt"

func RecursiveFactorial(nb int) int {
	factorial := 1
	if nb < 0 || nb > 50 {
		factorial = 0
	} else if nb > 1 {
		factorial = nb * RecursiveFactorial(nb-1)
	}
	return factorial
}

func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}
