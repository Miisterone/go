package main

func IterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		result = 0
	} else if power > 0 {
		result = nb * IterativePower(nb, power-1)
	}
	return result
}

func main() {
	println(IterativePower(4, 3))
}
