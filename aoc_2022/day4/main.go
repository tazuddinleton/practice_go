package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	from int
	to   int
}

func newRange(pair string) Range {
	r := strings.Split(pair, "-")
	from, _ := strconv.Atoi(r[0])
	to, _ := strconv.Atoi(r[1])
	return Range{from: from, to: to}
}

func (r Range) contains(other Range) bool {
	return r.from <= other.from && r.to >= other.to
}

func (r Range) overlaps(other Range) bool {
	return r.contains(other) || r.from <= other.from && other.from >= r.to
}

func main() {
	overlaps, contains := findOverlaps("input.txt")
	fmt.Println("Overlaps: ", overlaps, "Contains: ", contains)
}

func findOverlaps(input string) (overlaps int, contains int) {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		if err := sc.Err(); err != nil {
			panic(err)
		}

		line := sc.Text()
		pairs := strings.Split(line, ",")
		ranges := []Range{newRange(pairs[0]), newRange(pairs[1])}

		fmt.Println("Line: ", line)
		if ranges[0].contains(ranges[1]) || ranges[1].contains(ranges[0]) {
			fmt.Println("Found pair that contains other pair: ", ranges)
			contains++
		}

		if ranges[0].overlaps(ranges[1]) {
			fmt.Println("Found overlap: ", ranges)
			overlaps++
		}
	}
	return overlaps, contains
}
