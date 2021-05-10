package main

import (
	"fmt"
)

func Past(h, m, s int) int {
	milis := 0
	milim := 0
	milih := 0

	result := 0

	if 0 <= s && s <= 59 {
		milis = s * 1000
	}
	if 0 <= m && m <= 59 {
		milim = (m * 60) *1000
	}
	if 0 <= m && m <= 59 {
		milih = (h * 60 * 60) *1000
	}
	result = milis + milim + milih
	return result
}

func main()  {
	fmt.Println(Past(1, 1, 1))
}