package module

func Classic(a int) int {
	res := 1
	for i := 0; i < a; i++{
		res = res * a
	}
		return res
}

func While(a int) int {
	res := 1
	i := 0
	for i < a {
		res *= a
		i++
	}
	return res
}

func Range(str string, key int) string {

	var	result string

	for _, char := range str {
		temp := ConvertInt32ToInt(char)
		if temp + key < 'z' {
			result += ConvertIntToString(temp + key)
		} else {
			result += ConvertIntToString(temp - 26 + key)
		}
	}
	return result
}

func ConvertInt32ToInt(a int32) int {
	return int(a)
}

func ConvertIntToString(a int) string {
	return string(a)
}

func PrintLw(s string) string {
	var last_index int
	for i, c := range s {
		if c == ' ' {
			last_index = i + 1
		}
	}
	return s[last_index:]
}