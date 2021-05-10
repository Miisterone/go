package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			log.Fatal(err)
		}
	}

	for _, args := range os.Args[1:] {

		text, err := ioutil.ReadFile(args)
		if err != nil {
			fmt.Println("ERREUR :", err)
		}
		for _, result := range string(text) {
			z01.PrintRune(result)
		}
		z01.PrintRune('\n')
	}
}
