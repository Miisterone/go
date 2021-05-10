package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		z01.PrintRune('')
	}

}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
