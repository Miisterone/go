package student

func RecursivePower(nb int, power int) int {

	result := 1

	if power < 0 {
		return 0
	}

	if power > 0 {
		result = nb * RecursivePower(nb, power-1)
	}
	return result

}
