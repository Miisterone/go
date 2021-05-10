package Swap

func Swap(a *int, b *int) {
	c := *b
	*b = *a
	*a = c
}
