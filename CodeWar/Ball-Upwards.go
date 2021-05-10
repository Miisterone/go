package main

import (
	"fmt"
	"math"
)

func MaxBall(v0 int) int{
	var re float64
	var maxh float64
	var minh float64


	v := float64(v0) / 3.6
	g := 9.81
	time := 0.001
	re = v*time - 0.5*g*(time * time )
	var z float64
	count := 0


	minh = re

	for re >= 0 {
		count++
		time += 0.001
		re = v*time - 0.5*g*(time*time)
		minh = re
		if minh > maxh {
			maxh = minh
			z = time * 10
		}
	}
	return int(math.Round(z))
}

func main() {
	fmt.Println(MaxBall(45))
}
