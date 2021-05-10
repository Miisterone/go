package main

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	runes := [10]rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	v := 0
	for nb1 := 0; nb1 <= 9; nb1++ {
		for nb2 := 0; nb2 <= 9; nb2++ {
			for nb3 := 0; nb3 <= 9; nb3++ {
				for nb4 := 0; nb4 <= 9; nb4++ {
					if (nb1*10)+nb2 < (nb3*10)+nb4 {
						if v != 0 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
							v = 0
						}
						v++
						z01.PrintRune(runes[nb1])
						z01.PrintRune(runes[nb2])
						z01.PrintRune(' ')
						z01.PrintRune(runes[nb3])
						z01.PrintRune(runes[nb4])
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintComb2()
}
