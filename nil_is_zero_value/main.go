package main

import "fmt"

func main() {

}

func problem8() {
	/*
		nil is a predeclared identifier that represents
		the zero value of a pointer, slice, map, func, channel or interface
	*/

	a := []string{"A", "B", "C", "D", "E"}
	a = nil
	fmt.Println(a, len(a), cap(a))
}
