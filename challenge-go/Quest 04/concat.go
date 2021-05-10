package main

import (
	"fmt"
)

func Concat(str1 string, str2 string) string {
	str3 := str1 + str2
	return str3
}

func main() {
	fmt.Println(Concat("Hello!", " How are you?"))

}
