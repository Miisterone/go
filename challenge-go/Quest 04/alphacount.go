package main

import (
	"fmt"
)

func AlphaCount(s string) int {
	compteur := 0
	for _, nb := range s {
		if nb >= 'a' && nb <= 'z' || nb >= 'A' && nb <= 'Z' {
			compteur = compteur + 1
		}
	}
	return compteur
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
