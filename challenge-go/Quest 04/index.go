package main

func Index(s string, toFind string) int {
	sRune := []rune(s)
	tfRune := []rune(toFind)
	lensRune := 0
	lentfRune := 0
	indexResult := -1

	for range tfRune {
		lentfRune++
	}

	for range sRune {
		lensRune++
	}

	if lentfRune == 0 {
		return 0
	} else if lensRune == 0 {
		return -1
	}

	for index, val := range sRune {
		checkResult := 0
		if val == tfRune[0] && lensRune-index >= lentfRune {
			for j := 0; j < lentfRune; j++ {
				if sRune[j+index] == tfRune[j] {
					checkResult++
				}
			}
			if checkResult == lentfRune {
				return index
			}
		}
	}
	return indexResult
}

func main() {
	println(Index("Hello!", "l"))
	println(Index("Salut!", "alu"))
	println(Index("Ola!", "hOl"))
}
