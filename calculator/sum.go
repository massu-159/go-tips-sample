package calculator

import "fmt"

var offset float64 = 1
var Offset float64 = 1

func Sum(x, y float64) float64 {
	fmt.Println("multiply", multiply(x, y))
	fmt.Println("Multiply", Multiply(x, y))
	return x + y + offset
}