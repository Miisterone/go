package student

func Split(s, sep string) []string {

	var result []string
	separation := len(sep)
	var check int

	for i := range s {
		if i < len(s)-separation-1 {
			if s[i:i+separation] == sep {
				result = append(result, s[check:i])
				check = i + separation
			}
		}
	}
	return append(result, s[check:])
}
