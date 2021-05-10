package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {

	stringNumber  := strings.Split(in, " ")

	if len(stringNumber) < 2 {
		return in + " " + in
	}

	max, _ := strconv.Atoi(stringNumber[0])
	min, _ := strconv.Atoi(stringNumber[0])

	for _, a := range stringNumber {

		n, _ := strconv.Atoi(a)
		if max < n {
			max = n
		}
		if min > n {
			min = n
		}
	}
	return strconv.Itoa(max) + " " + strconv.Itoa(min)

}

func main() {
	fmt.Println(HighAndLow("1 1 1"))
	fmt.Println(HighAndLow("1 2 -3 4 5"))
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}
