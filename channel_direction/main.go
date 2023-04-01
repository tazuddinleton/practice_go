package main

import "fmt"

func main() {
	
}

func problem4() {
	// Channels have direction. It can be bidirectional and unidirectional.
	// This is a send-only unidirectional channel
	sendOnly := make(chan<- string)

	// This is a receive-only unidirectional channel
	receiveOnly := make(<-chan string)

	var msg string
	select {
	case r := <-receiveOnly:
		fmt.Println("Receive only unidirectional channel")
		msg = r
	case sendOnly <- msg:
		fmt.Println("send-only unidirectional channel")
	default:
		fmt.Println("default case")
	}
}
