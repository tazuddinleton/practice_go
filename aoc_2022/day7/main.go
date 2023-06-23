package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	root := readInput("./input.txt")
	printTree(root, 0)
}

func readInput(path string) *File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)

	root := newFile("/", true, 0)
	_ = sc.Scan() // skip the first line
	curr := root
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			panic(err)
		}

		line := strings.Fields(sc.Text())

		if isChangeDir(line) {
			fmt.Println("changing dir to ", line[2])
			if line[2] == ".." {
				curr = curr.getParent()
			} else {
				curr = curr.getChild(line[2])
			}
		} else if isListFiles(line) {
			continue
		} else if isDir(line) {
			fmt.Println("creeating new dir: ", line[1])
			curr.addNew(newFile(line[1], true, 0))
		} else {
			fmt.Println("creeating new file: ", line[1])
			size, err := strconv.Atoi(line[0])
			if err != nil {
				panic(err)
			}
			curr.addNew(newFile(line[1], false, size))
		}
	}

	return root
}

func printTree(file *File, level int) {
	level++
	tab := strings.Repeat("  ", level)
	if file.IsDir {
		fmt.Println(fmt.Sprintf("%s- %s", tab, file.Name))
	} else {
		fmt.Println(fmt.Sprintf("%s- %s %d", tab, file.Name, file.Size))
	}
	ch := file.getChildren()
	for i := 0; i < len(ch); i++ {
		printTree(ch[i], level)
	}
}

func isListFiles(line []string) bool {
	return line[0] == "$" && line[1] == "ls"
}

func isChangeDir(line []string) bool {
	return line[0] == "$" && line[1] == "cd"
}

func isDir(line []string) bool {
	return line[0] == "dir"
}

func isRoot(line []string) bool {
	return line[0] == "$" && line[1] == "cd" && line[2] == "/"
}
