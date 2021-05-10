package main

import (
	"fmt"

	"./student"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Println(student.Split(s, "HA"))
}
