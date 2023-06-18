package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	readInput("./input.txt")
}

func readInput(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)

	var root File
	var curr File
	curr = root
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			panic(err)
		}

		line := strings.Fields(sc.Text())
		if line[0] == "$" && line[1] == "cd" {
			curr.addNew(newFile(line[2], true, 0))
		}
		// fmt.Println(line)
	}
}
