package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(N int) {
	var runes []int // tableau de rune
	if N == 0 {
		z01.PrintRune(rune(48)) // Afficher 0
		return
	} else {
		for N > 0 {
			runes = append(runes, N%10) // Ajouter une valeur à un tableau
			Reste := N % 10
			N = (N - Reste) / 10 // Transformer le int en tableau de chiffre
		}
	}

	verif := 0 // si n != 0 pour rentrer dans la boucle

	for verif != 1 {
		verif = 0 // rester dans la boucle
		for index := 0; index < len(runes); index++ {
			if index == len(runes)-1 {
				verif = 1
				break
			}
			if runes[index] > runes[index+1] { // Switch des char
				runes[index], runes[index+1] = runes[index+1], runes[index]
				break
			}
		}
	}
	for index := 0; index < len(runes); index++ { // Tant que index est inférieur au nombre de char dans le taleau on affiche dans l'ordre
		z01.PrintRune(rune(runes[index] + 48))
	}
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
