package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func problem2(result []int) []int {
	m := make(map[int]int)
	var output []int

	for i := 0; i < len(result); i++ {
		if _, ok := m[result[i]]; ok {
			output = append(output, result[i])
		}
		m[i] = result[i]
	}
	// m: 1,2,3,4

	for j := 1; j <= len(result); j++ {
		if _, ok := m[j]; ok {
			continue
		}
		output = append(output, j)
	}
	return output
}

func problem1() {
	digits := "0123456789"
	num := "210"

	m := make(map[int]int)
	for i, d := range digits {
		fmt.Println(d)

		k, _ := strconv.Atoi(string(d))
		m[i] = k
	}

	fmt.Println(m)
	sum := 1
	for _, n := range num {
		fmt.Println(n)
		l, _ := strconv.Atoi(string(n))
		sum += m[l]
	}

}
