package main

import "fmt"

func QuarterOf(month int) int {
	if month >= 1 && month <= 3 {
		return 1
	}
	if month >= 4 && month <= 6 {
		return 2
	}else if month >= 7 && month <= 9 {
		return 3
	}
	return 4
}

func main()  {
	fmt.Println(QuarterOf(3))
}