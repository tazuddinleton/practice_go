package main

import (
	"fmt"
	"sync"
)

func main() {
	s := "a string"
	switch s {
	case "a string":
		fmt.Println("it is a string")
	default:
		fmt.Println("default for switch")
	}

	wg := sync.WaitGroup{}
	wg.Add(10)
	c := broadCast(10, &wg)

	//select {
	//case v := <-c:
	//	fmt.Println("i: ", v)
	//}
	for i := range c {
		fmt.Println("i: ", i)
	}

	wg.Wait()
}

func broadCast(n int, wg *sync.WaitGroup) <-chan int {
	c := make(chan int)
	//defer func() { close(c) }()
	for i := 0; i < n; i++ {
		go func(v int) {
			c <- v
			wg.Done()
		}(i)
	}
	return c
}
