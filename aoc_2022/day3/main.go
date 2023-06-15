package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	// sum := getSum("./input.txt")
	sum := badgePrioritySum("./input.txt")
	fmt.Println("Sum: ", sum)
}

func getSum(input string) int {
	f, err := os.Open(input)
	if err != nil {
		log.Panic(err)
	}

	sc := bufio.NewScanner(f)

	sum := 0

	for sc.Scan() {
		if err := sc.Err(); err != nil {
			log.Panic(err)
		}

		line := sc.Text()
		fmt.Println("Rucksack: ", line)
		runes := []rune(line)
		comp1 := runes[0 : len(runes)/2]
		comp2 := runes[len(runes)/2:]

		fmt.Println("comp1 : ", string(comp1))
		fmt.Println("comp2 : ", string(comp2))

		c, _ := getCommon(comp1, comp2)
		sum += priority(c)
	}
	return sum
}

func badgePrioritySum(input string) int {
	f, err := os.Open(input)
	if err != nil {
		log.Panic(err)
	}

	sc := bufio.NewScanner(f)

	sum := 0
	count := 0

	groups := []string{}
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			log.Panic(err)
		}

		count++
		line := sc.Text()
		fmt.Println("Rucksack: ", line)

		groups = append(groups, line)
		if count%3 == 0 {
			fmt.Println("Group: ", count/3)
			c, err := getCommon2(groups)

			if err != nil {
				panic(err)
			}
			sum += priority(c)
			groups = groups[:0]
		}
	}
	return sum
}

func getCommon2(group []string) (rune, error) {
	dic := make(map[rune]bool)
	dic2 := make(map[rune]bool)
	for _, r := range []rune(group[0]) {
		if _, ok := dic[r]; !ok {
			dic[r] = true
		}
	}
	for _, r := range []rune(group[1]) {
		if _, ok := dic2[r]; !ok {
			dic2[r] = true
		}
	}
	for _, r := range []rune(group[2]) {
		_, ok := dic[r]
		_, ok2 := dic2[r]
		if ok && ok2 {
			fmt.Println("Common: ", string(r))
			return r, nil
		}
	}
	return rune(0), errors.New("not found")
}

func getCommon(comp1, comp2 []rune) (rune, error) {
	frequency := make(map[rune]bool)
	for _, r := range comp1 {
		if _, ok := frequency[r]; !ok {
			frequency[r] = true
		}
	}
	for _, r := range comp2 {
		if _, ok := frequency[r]; ok {
			fmt.Println("Common item: ", string(r))
			return r, nil
		}
	}
	return rune(0), errors.New("not found")
}

func priority(r rune) int {
	priority := 0
	if r > 97 {
		priority = int(r) - 96
	} else {
		priority = int(r) - 38
	}

	fmt.Println("Priority: ", priority)
	return priority
}
