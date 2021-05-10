package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	erreur := "not possible..."
	erreur2 := "incomplete..."
	args := os.Args[1:]

	count := 0

	if len(args) <= 1 {
		fmt.Println(erreur2)
	}

	if len(args) == 2 {
		x, err := strconv.Atoi(os.Args[1])
		y, err2 := strconv.Atoi(os.Args[2])

		if x == 0 || y == 0 {
			fmt.Println(erreur)
		}

		if err == nil && err2 == nil {
			for h := 0; h < y; h++ {
				if h == 0 {
					if x == 1 {
						fmt.Print("X")
					} else{
						fmt.Print("X")
						for l := 0; l < x-2; l++ {
							fmt.Print(string(alphabet[count]))
							count++
							if count == 26 {
								count = 0
							}
						}
						fmt.Print("X")
					}
				} else if h == y-1 {
					fmt.Print("\n")
					if x == 1 {
						fmt.Print("X")
					} else {
						fmt.Print("X")
						for l := 0; l < x-2; l++ {
							fmt.Print(string(alphabet[count]))
							count++
							if count == 26 {
								count = 0
							}
						}
						fmt.Print("X")
					}
				} else {
					fmt.Print("\n")
					if x == 1 {
						fmt.Print("|")
					}else {
						fmt.Print("|")
						for l := 0; l < x-2; l++ {
							fmt.Print(" ")
						}
						fmt.Print("|")
					}
				}
			}
		}
	}
}
