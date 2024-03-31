package main

import (
	"fmt"
	"math"
)


func triangleArea(a, b, c float64) float64 {
	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return area
}

func main() {
	var a, b, c float64


	a, b, c = 3, 4, 5

	area := triangleArea(a, b, c)
	fmt.Printf("The area of the triangle with sides %.2f, %.2f, and %.2f is %.2f\n", a, b, c, area)
}
