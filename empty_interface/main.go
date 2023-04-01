package main

import "fmt"

type S struct {
}

func (s S) F() {

}

type IF interface {
	F()
}

func InitType() S {
	var s S
	return s
}

func InitPointer() *S {
	var s *S
	fmt.Printf("Init Pointer %v %t \n", s, s)
	return s
}

func InitEfaceType() interface{} {
	var s S
	fmt.Printf("InitEfaceType %v %t \n", s, s)
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	fmt.Printf("InitEfacePointer %v %t \n", s, s)
	return s
}

func InitIfaceType() IF {
	var s S
	fmt.Printf("InitIfaceType %v %t\n", s, s)
	return s
}

func InitIfacePointer() IF {
	var s S
	fmt.Printf("InitIfacePointer %v %t\n", s, s)
	return s
}

func main() {
	println(InitPointer() == nil)
	println(InitEfaceType() == nil)
	println(InitEfacePointer() == nil)
	println(InitIfaceType() == nil)
	println(InitIfacePointer() == nil)
}
