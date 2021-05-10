package main

import (
	"fmt"
)

func IsUpper(s string) bool {

	for _, lettre := range s {
		if lettre < 65 || lettre > 90 { // Table Ascii
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

}
