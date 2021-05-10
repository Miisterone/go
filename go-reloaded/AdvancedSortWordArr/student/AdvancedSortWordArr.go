package student

func Compare(a string, b string) int {

	if a > b {
		return 1
	}
	return 0
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) {

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if f(a[j], a[j+1]) == 1 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
