package main

import (
	"fmt"
)

func IsLower(s string) bool {
	for _, lettre := range s {
		if lettre < 97 || lettre > 122 { // Table Ascii
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}
