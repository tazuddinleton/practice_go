package main

import "fmt"

// https://stackoverflow.com/questions/20254834/x-stringsat-sun-vs-x-stringsat-sun
func main() {

	variadic(1, "hello", 1, 2, 3, 4, 5)
	argumentsToVariadicFunction()
	arrayLiteral()
}

// Variadic function parameters
func variadic(a int, b string, c ...int) {

	fmt.Println("a: ", a, "b: ", b, "Variadic parameter c: ", c)
}

func argumentsToVariadicFunction() {
	args := []int{1, 2, 3, 4}
	variadic(1, "argumentsToVariadicFunction", args...)
}

func arrayLiteral() {
	a := []int{1}
	b := [...]int{1, 2, 3, 4}
	fmt.Println("len(a)", len(a), "len(b)", len(b))
}

/*
	Another place is in the go command
	go test ./...
*/
