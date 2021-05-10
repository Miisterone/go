package main

import "fmt"

func Hero(bullets,dragons int) bool{
	return bullets < dragons * 2 ==false
}
func main()  {
	fmt.Println(Hero(100, 40))
}