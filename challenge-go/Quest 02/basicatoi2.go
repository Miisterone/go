package main

import "fmt"

func BasicAtoi2(s string) int {
	total := 0

	for _, char := range s { // pour, tout le temps, en fonction du rang du charactère ( rang + 1 )
		intChat := int(char)
		if intChat >= 48 && intChat <= 57 {

			total = total*10 + (intChat - 48) // Total : Charactère final *10 pour le changer de place, et le reste on prend le charactère et on fait -48 (le 0 en utf8) pour le passer en chiffre)
		} else {
			return 0
		}
	}
	return total
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
