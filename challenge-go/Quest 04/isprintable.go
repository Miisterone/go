package main

import (
	"fmt"
)

func IsPrintable(s string) bool {
	for _, lettre := range s {
		if lettre < 32 || lettre > 126 { // Table Ascii
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}
