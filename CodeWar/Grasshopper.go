package main

func CheckForFactor(base int, factor int) bool {
	result := 0

	result = base % factor
	if result != 0 {
		return false
	}else {
		return true
	}
}