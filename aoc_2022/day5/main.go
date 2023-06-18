package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readInput("./input.txt")
}

func readInput(input string) {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)
	lineNum := 1
	lines := []string{}
	for sc.Scan() && lineNum != 9 {
		line := sc.Text()
		fmt.Println("Line: ", line)
		lines = append(lines, line)
		lineNum++
	}

	stacks := makeStacks(lines)
	PrintStacks(stacks[:])

	for sc.Scan() {
		move := strings.Fields(sc.Text())
		if len(move) == 0 {
			continue
		}
		count, err := strconv.Atoi(move[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(move[3])
		if err != nil {
			log.Panic(err)
		}
		to, err := strconv.Atoi(move[5])
		if err != nil {
			log.Panic(err)
		}

		fmt.Println("move", count, "from", from, "to", to)
		moveCratesUsing9001(stacks[:], count, from, to)
		PrintStacks(stacks[:])
	}

	PrintStacks(stacks[:])
	PrintTopCrates(stacks[:])
}

func moveCrates(stacks []Stack, howMany, from, to int) {
	for i := 0; i < howMany; {
		crate := stacks[from-1].Pop()
		stacks[to-1].Push(crate)
		i++
	}
}

func moveCratesUsing9001(stacks []Stack, count, from, to int) {
	crate := stacks[from-1].PopMany(count)
	fmt.Println("moving crates: ", crate, "from", from, "to", to)
	stacks[to-1].PushMany(crate)
}

// [T] [V]                     [W]
// [V] [C] [P] [D]             [B]
// [J] [P] [R] [N] [B]         [Z]
// [W] [Q] [D] [M] [T]     [L] [T]
// [N] [J] [H] [B] [P] [T] [P] [L]
// [R] [D] [F] [P] [R] [P] [R] [S] [G]
// [M] [W] [J] [R] [V] [B] [J] [C] [S]
// [S] [B] [B] [F] [H] [C] [B] [N] [L]
//
//	1   2   3   4   5   6   7   8   9
func makeStacks(lines []string) [9]Stack {
	stacks := [9]Stack{{}, {}, {}, {}, {}, {}, {}, {}, {}}

	log.Println("lines ", lines)
	stackIndex := 0
	for i := len(lines) - 1; i > -1; i-- {
		line := lines[i]
		fmt.Println("line: ", line, "line num: ", i)
		for j := 1; j < len(line)-1; j += 4 {
			crate := strings.TrimSpace(string(line[j]))
			if crate == "" {
				stackIndex++
				continue
			}
			stacks[stackIndex].Push(crate)
			stackIndex++
		}
		stackIndex = 0
	}
	return stacks
}
