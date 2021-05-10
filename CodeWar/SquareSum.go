package main

import "fmt"

func SquareSum(numbers []int) int{
	result := 0

	for _, a := range numbers {
		result += a * a
	}
	return result
}
func main() {
	fmt.Println(SquareSum([]int{1,2}))
}
