package student

func Atoi(s string) int {

	total := 0

	signe := false

	for i, char := range s {
		intChar := int(char)
		if intChar == '+' {
			if i != 0 {
				return 0
			}
		} else if intChar == '-' {
			if i != 0 {
				return 0
			} else if i == 0 {
				signe = true
			}

		} else if intChar >= '0' && intChar <= '9' {

			total = total*10 + (intChar - 48)

		} else {
			return 0
		}
	}
	if signe == true {
		total *= -1
	}
	return total
}
