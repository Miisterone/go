package main

import "fmt"

func Summation(n int) int {
	i := 0
	sum := 0
	add := 0
	for i != n {
		i++
		add += 1
		sum += add
	}
	return sum
}

func main()  {
	fmt.Println(Summation(8))
}