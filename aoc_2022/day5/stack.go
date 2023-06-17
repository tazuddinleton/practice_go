package main

import "fmt"

type Stack struct {
	data []string
}

func (s *Stack) Push(item string) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() string {
	item := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return item
}

func PrintStacks(stacks []Stack) {
	for _, s := range stacks {
		fmt.Println(s.data, len(s.data))
	}
}

func PrintTopCrates(stacks []Stack) {
	for _, s := range stacks {
		fmt.Print(s.data[len(s.data)-1])
	}
	fmt.Print("\n")
}
