package main

import (
	"fmt"
)

func MakeUpperCase(str string) string {

	result := ""

	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result += string(char - 32) // x += y  fait x = x + y
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(MakeUpperCase("hello"))
	fmt.Println(MakeUpperCase("hello world"))
	fmt.Println(MakeUpperCase("hello world !"))
	fmt.Println(MakeUpperCase("heLlO wORLd !"))
	fmt.Println(MakeUpperCase("1,2,3 hello world!"))
}
