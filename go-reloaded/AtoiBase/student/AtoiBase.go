package student

func AtoiBase(s string, base string) int {

	Check := true
	result := 0

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

	if Check == true {

		for i, char := range s {
			index := index(string(char), base)
			power := puissance(len(base), len(s)-(i+1)) //len(s)-(i+1) Pour crée les dizaines,centaines,etc
			result = result + power*index               //(index du caractère dans la base) * (taille de la base)   ^ (rang du caractère dans le nombre de départ)
		}
	} else {
		return 0
	}
	return result

}

func index(s string, base string) int {

	for i, c := range base {
		if string(c) == s {
			return i
		}
	}
	return ' '
}

func puissance(nbr int, power int) int {

	result := nbr

	if power == 0 {
		return 1

	}
	for i := 1; i < power; i++ {
		result = result * nbr
	}
	return result
}
