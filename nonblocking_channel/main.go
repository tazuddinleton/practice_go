package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

func main() {

}

func problem3() {
	// Non-blocking channels
	/*
		In Go basic channels are blocking. We can create non-blocking channels by
		using select with default case
	*/

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("message received", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("message sent", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case sig := <-signals:
		fmt.Println("signal received", sig)
	case msg := <-messages:
		fmt.Println("message received", msg)
	default:
		fmt.Println("no activity")
	}

}

func problem5() {

	wg := sync.WaitGroup{}
	wg.Add(2)
	var service1 func(chan string)
	var service2 func(chan string)
	service1 = func(c chan string) {
		log.Println("hello from service 1")
		c <- "hello from service 1"
		defer wg.Done()
	}

	service2 = func(c chan string) {
		log.Println("hello from service 2")
		c <- "hello from service 2"
		defer wg.Done()
	}

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	log.Println("num gorotine", runtime.NumGoroutine())
	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res)
	case res := <-chan2:
		fmt.Println("Response from service 2", res)
		//default:
		//	fmt.Println("No response received")
	}
	log.Println("num gorotine", runtime.NumGoroutine())
}
