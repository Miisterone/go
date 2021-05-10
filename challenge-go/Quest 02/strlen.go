package StrLen

func StrLen(s string) int { //Crée un string s (sa longueur est de 1 car il y a que "s",ex : "Hello"= longueur de 5)
	compteur := 0
	for range s { //Parcours la longueur de string s
		compteur = compteur + 1
	}
	return compteur
}
