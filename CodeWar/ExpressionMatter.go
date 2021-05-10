/*
Étant donné trois entiers a ,b,c, reviennent le plus grand nombre obtenu après l' insertion des opérateurs et des supports suivants :+,*,()
En d'autres termes, essayez toutes les combinaisons de a, b, c avec [* + ()] et renvoyez le maximum obtenu
 */

package main

import (
	"fmt"
)

func EM(a int, b int, c int) int {
	MaxResult := 0
	LastResult := 0

	MaxResult = a * (b + c)
	LastResult = a * b * c
	if LastResult > MaxResult {
		MaxResult = LastResult
		LastResult = 0
	}else {
		LastResult = a + b * c
	}
	if LastResult > MaxResult {
		MaxResult = LastResult
		LastResult = 0
	}else {
		LastResult = (a + b) * c
	}
	if LastResult > MaxResult {
		MaxResult = LastResult
		LastResult = 0
	}else {
		LastResult = a + b + c
	}
	if LastResult > MaxResult {
		MaxResult = LastResult
		LastResult = 0
	}
	return MaxResult

}

func main() {
	fmt.Println(EM(2, 2, 2))
}


