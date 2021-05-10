package main

import "fmt"

func Solution(word string) string {
	result := ""
	for _, a := range word { // range permet de parcourir le string
		result = string(a) + result
	}
	return result
}

func main() {
	fmt.Println(Solution("World"))
}
