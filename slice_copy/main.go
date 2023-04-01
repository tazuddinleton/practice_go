package main

import "fmt"

func main() {
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6, 7}
	n2 := copy(s2, s3)

	fmt.Printf("n2=%d, s2=%v, s3=%v\n", n2, s2, s3)
}
