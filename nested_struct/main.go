package main

import "fmt"

type Query struct {
	Page int
}

type Payload struct {
	Query Query
	Outer int
}

func main() {
	var payload *Payload
	payload = &Payload{
		Query: Query{Page: 1},
		Outer: 1,
	}

	fmt.Printf("initial value: %d outer: %d\n", payload.Query.Page, payload.Outer)
	p := *payload
	p.Query.Page = 10
	p.Outer = 100
	fmt.Printf("after change value: %d outer: %d\n", payload.Query.Page, payload.Outer)
}
