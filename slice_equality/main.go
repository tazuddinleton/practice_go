package main

import (
	"fmt"
	"reflect"
)

func main() {
	problem6()
}

func problem6() {
	// comparing []int with []any gives false
	s1 := []int{1, 2, 3}
	s2 := []any{1, 2, 3}
	// Invalid operation exception
	// fmt.Println(s1 == s2)
	eq := reflect.DeepEqual(s1, s2)
	fmt.Println("Are the equal?", eq)
}
