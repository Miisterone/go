package main

import (
	"fmt"
)

func ToLower(s string) string {
	result := ""
	for _, val := range s {
		if val <= 90 && val >= 65 {
			result += string(val + 32)
		} else {
			result += string(val)
		}
	}
	return result
}

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}
