package main

import (
	"fmt"
)

func TrimAtoi(s string) int {
	result := 0
	compteur := 0
	négatif := 0
	for _, char := range s {
		if char >= 48 && char <= 57 { // cherche des chiffres
			result = result*10 + (int(char) - 48) //Pour le transformer en int
			compteur = compteur + 1
		} else if char == 45 { // si il est négatif
			if compteur < 1 {
				négatif = 1
			}
		}

	}
	if négatif == 1 {
		result = result * -1
	}
	return result
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
}
