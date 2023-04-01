package main

import "fmt"

func main() {

	fmt.Println("main() started")
	go service()
	select {}

	fmt.Println("main() stopped")
}

func service() {
	fmt.Println("Hello from service")
}
