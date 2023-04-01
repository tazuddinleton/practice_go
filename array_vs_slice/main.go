package main

import "fmt"

func main() {

}

/*
	You can make the slice nil but keep the memory allocation
	using a[:0] expression.

	Only difference between slice and array is: array is declared with fixed size and slice isn't.
	[5]int is an array []int is a slice.

*/

func problem16() {
	// Guess the output
	a := []string{"A", "B", "C", "D", "E"}
	a = a[:0]
	fmt.Println(a, len(a), cap(a))

}
