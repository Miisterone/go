package main

import (
	"fmt"
)

func IsNumeric(s string) bool {
	for _, lettre := range s {
		if lettre < 48 || lettre > 57 { // Table Ascii
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}
