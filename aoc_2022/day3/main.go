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
	sum := badgePrioritySum("./groups.txt")
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

	groups := []rune{}
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			log.Panic(err)
		}

		count++
		line := sc.Text()
		fmt.Println("Rucksack: ", line)
		runes := []rune(line)
		groups = append(groups, runes...)

		if count%3 == 0 {
			fmt.Println("Group: ", count/3)
			c, _ := getCommon(groups[0:len(groups)/2], groups[len(groups)/2:])
			sum += priority(c)
			groups = groups[:0]
		}
	}
	return sum
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
