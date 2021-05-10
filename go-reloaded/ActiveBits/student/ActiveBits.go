package student

func ActiveBits(n int) uint {

	reste := 0

	for n != 0 {
		if n%2 == 1 {
			reste++
		}
		n = n / 2
	}
	return uint(reste)
}
