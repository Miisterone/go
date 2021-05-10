package main

import "fmt"

func IterativeFactorial(nb int) int {
	resultat := 1
	if nb < 0 || nb > 50 {
		return 0
	} else if nb == 0 {
		return 1

	} else {

		for i := 1; i <= nb; i++ {
			resultat = resultat * i
		}
	}
	if resultat < 0 {
		return 0
	}
	return resultat
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
