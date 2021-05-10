package main

import "fmt"

func Greet(name string) string {

	result := "Hello, " + name + " how are you doing today?"

	return result
}

func main() {
	fmt.Println(Greet("Ryan"))
}