package main

import (
	"fmt"

	"./student"
)

func main() {

	fmt.Println(student.Atoi("12345"))
	fmt.Println(student.Atoi("0000000012345"))
	fmt.Println(student.Atoi("012 345"))
	fmt.Println(student.Atoi("Hello World!"))
	fmt.Println(student.Atoi("+1234"))
	fmt.Println(student.Atoi("-1234"))
	fmt.Println(student.Atoi("++1234"))
	fmt.Println(student.Atoi("--1234"))

}
