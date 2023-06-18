package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	data []string
}

func (s *Stack) Push(item string) {
	s.data = append(s.data, item)
}

func (s *Stack) PushMany(item string) {
	s.data = append(s.data, strings.Split(item, "")...)
}

func (s *Stack) Pop() string {
	item := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return item
}

func (s *Stack) PopMany(count int) string {
	item := s.data[len(s.data)-count:]
	s.data = s.data[0 : len(s.data)-count]
	return strings.Join(item, "")
}

func PrintStacks(stacks []Stack) {
	for i, s := range stacks {
		fmt.Println(i+1, s.data, len(s.data))
	}
}

func PrintTopCrates(stacks []Stack) {
	for _, s := range stacks {
		fmt.Print(s.data[len(s.data)-1])
	}
	fmt.Print("\n")
}
