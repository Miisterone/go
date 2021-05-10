package main

import (
	"fmt"
)

func ConcatParams(args []string) string {

	s := "" // Création d'un string vide

	for i, arg := range args { //Parcour tous le string args
		s = s + arg          // Remplace le string vide et ajout "Hello", "how", "are", "you?" ou autre
		if len(args)-1 > i { //  if len(args)-1 > i = si l'avant dernier argument (car dernier - 1 = avant dernier) est supérieur à i
			s = s + "\n" //Reviens a la ligne
		}
	}
	return s

}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
