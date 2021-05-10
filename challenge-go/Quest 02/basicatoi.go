package main

import "fmt"

func BasicAtoi(s string) int {
	total := 0

	for _, char := range s { // pour, tout le temps, en fonction du rang du charactère ( rang + 1 )
		total = total*10 + (int(char) - 48) // Total : Charactère final *10 pour le changer de place, et le reste on prend le charactère et on fait -48 (le 0 en utf8) pour le passer en chiffre)
	}
	return total
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
