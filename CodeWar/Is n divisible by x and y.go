package main

import "fmt"

func IsDivisible(n, x, y int) bool {
	if n % x == 0 {
		if n % y == 0 {
			return true
		}
	}
	return false
}

func main()  {
	 fmt.Println(IsDivisible(12,2,6))
}