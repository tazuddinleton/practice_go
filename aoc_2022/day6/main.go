package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	count, err := processSignal("./input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("count: ", count)
}

func processSignal(input string) (int, error) {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		err = sc.Err()
		if err != nil {
			panic(err)
		}

		line := sc.Text()

		fmt.Println("line: ", line)
		count, err := detectMark(line)
		if err != nil {
			panic(err)
		}
		return count, nil
	}
	return 0, errors.New("Failed to detect any signal")
}

func detectMark(signal string) (int, error) {
	count := 0
	q := &Queue{maxSize: 14}
	for _, c := range signal {
		q.Enque(c)
		count++
		if q.AllDifferent() {
			return count, nil
		}
	}
	return 0, errors.New("No mark detected!")
}
