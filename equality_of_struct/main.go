package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var s Shape
	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}

	// Structures are equal if their fields have identical values
	// You can compare using == operator or reflect.DeepEqual function
	fmt.Println(s == r)
}
