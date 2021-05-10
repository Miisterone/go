package main

import (
	"fmt"
)

func ToUpper(s string) string {
	result := ""
	for _, val := range s {
		if val <= 122 && val >= 97 {
			result += string(val - 32)
		} else {
			result += string(val)
		}
	}
	return result
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
