package main

import (
	"fmt"

	"./student"
)

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(student.RecursivePower(arg1, arg2))
}
