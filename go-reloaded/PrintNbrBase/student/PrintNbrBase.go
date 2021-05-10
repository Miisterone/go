package student

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {

	Check := true

	if len(base) < 2 { //A base must contain at least 2 characters
		Check = false
	}

	for i := 0; i < len(base); i++ { //A base should not contain + or - characters.
		if base[i] == '+' || base[i] == '-' {
			Check = false
		}
		for i2 := i + 1; i2 < len(base); i2++ { //Each character of a base must be unique.
			if base[i] == base[i2] {
				Check = false
			}
		}
	}

	var list []rune

	if Check == true {
		if nbr < 0 {
			z01.PrintRune('-')
			nbr = nbr * -1
		}
		for nbr > 0 {

			temp1 := nbr % len(base)
			nbr = nbr / len(base)
			list = append(list, rune(base[temp1]))

		}
		for i := len(list) - 1; i >= 0; i-- {
			z01.PrintRune(list[i])
		}

		if nbr == -9223372036854775808 {
			for _, a := range "9223372036854775808" {
				z01.PrintRune(a)
			}
		}
	} else { //If the base is not valid
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}
