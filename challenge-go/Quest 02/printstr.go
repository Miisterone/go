package PrintStr

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, a := range s { // Parcour tous le string s et remplace par un a
		z01.PrintRune(a) //Affiche la valeur uft8
	}
}
