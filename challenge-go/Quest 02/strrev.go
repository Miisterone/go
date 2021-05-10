package StrRev

func StrRev(s string) (result string) {
	for _, v := range s { // parcour tous le string s et remplace par un v // _ est une variable qui sert a rien
		result = string(v) + result
	}
	return result //renvoie la valeur a l'endroit au elle a était
}
