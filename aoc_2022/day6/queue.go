package main

import "fmt"

type Queue struct {
	maxSize int
	data    []rune
}

func (q *Queue) Enque(item rune) {
	fmt.Println("Enque: ", item)
	if len(q.data) == q.maxSize {
		_ = q.Deque()
	}
	q.data = append(q.data, item)
	fmt.Println("Enque complete: ", q)
}

func (q *Queue) Deque() rune {
	item := q.data[0]
	fmt.Println("Deque: ", item)
	q.data = q.data[1:]
	return item
}

func (q Queue) AllDifferent() bool {
	dic := make(map[rune]int)
	for _, c := range q.data {
		if _, ok := dic[c]; !ok {
			dic[c] = 1
		} else {
			dic[c]++
		}
	}
	return len(dic) == q.maxSize
}
