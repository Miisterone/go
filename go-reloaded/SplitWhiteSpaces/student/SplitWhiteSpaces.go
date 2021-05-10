package student

func SplitWhiteSpaces(s string) []string {

	var result []string
	var mot string

	for _, char := range s {
		if char == ' ' {
			result = append(result, mot)
			mot = ""
		} else if char != ' ' {
			mot = mot + string(char)
		}
	}
	if mot != "" {
		result = append(result, mot)
	}
	return result

}
