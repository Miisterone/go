package main

import (
	"github.com/01-edu/z01"
)

func FirstRune(s string) rune {
	return []rune(s)[0] // [] c'est pour précisé que c'est un tableau
}

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
