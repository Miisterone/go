package main

import "fmt"

func Fibonacci(index int) int {
	F0 := 0
	F1 := 1
	if index < 0 {
		return -1
	}
	if index == 0 {
		return F0
	}
	if index == 1 {
		return F1
	}
	return Fibonacci(index-2) + Fibonacci(index-1)
}

func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}
