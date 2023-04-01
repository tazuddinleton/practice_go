package main

import "fmt"

func main() {

}

func problem6() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	// Invalid operation exception
	//fmt.Println(s1 == s2)
	fmt.Println(s1 == nil && s2 == nil)
}
