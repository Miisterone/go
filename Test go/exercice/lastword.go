package exercice

func LastW(s string) string {

	var lastword int

	for i, c := range s {
		if c == ' ' {
			lastword = i + 1
		}
	}
	return s[lastword:]

}
