package main

import "fmt"

func main() {
	//a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s := a[2:4]
	//
	//fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	//newS := append(s, 55, 66)
	//fmt.Printf("len=%d, cap=%d\n", len(newS), cap(newS))
	//
	//fmt.Println(s, newS)

	// len=9, cap=9
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("len=%d, cap=%d\n", len(a), cap(a))
	// len=1, cap=9, because the first element of the slice is 1 and last element of array is 9
	// if we count from 1..9 we get 9 element
	s := a[:1]
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	// len=0, cap=9, zero length but capacity still is 9
	s = a[:0]
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	// len=1, cap=1, because the first element of slice is 9 and last element of array is also 9
	s = a[8:]
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}
