package main

import (
	"fmt"

	"./student"
)

func main() {

	nbits := student.ActiveBits(7)
	fmt.Println(nbits)

}
