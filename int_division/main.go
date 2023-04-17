package main

import (
	"fmt"
	"math"
)

func main() {
	num1 := 21
	num2 := 10

	num3 := math.Ceil(float64(num1) / float64(num2))
	fmt.Println(num3)
}
