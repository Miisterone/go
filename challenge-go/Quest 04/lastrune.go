package main

import (
	"github.com/01-edu/z01"
)

func LastRune(s string) rune {
	return []rune(s)[len(s)-1] // len prend la longueur du string
}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}