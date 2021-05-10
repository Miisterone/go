package main

import (
	"fmt"
)

func IsAlpha(s string) bool {
	for _, lettre := range s {
		if lettre < 48 || lettre > 57 || lettre == 27 { // Table Ascii
			if lettre < 65 || lettre > 90 {
				if lettre < 97 || lettre > 122 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}
