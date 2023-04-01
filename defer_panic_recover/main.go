package main

import "fmt"

func main() {

	//a()
	//b()
	//fmt.Println(c())
	f()
	fmt.Println("returned normally from f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()
	fmt.Println("calling g")
	g(0)
}

func g(i int) {
	if i > 3 {
		fmt.Println("panicinig!")
		panic(i)
	}
	fmt.Println("incrementing...", i)
	g(i + 1)
}

func a() {
	i := 0
	defer fmt.Println(i)
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Printf("%d", i)
	}
}

func c() (i int) {
	defer func() { i++ }() // modifying returning value
	return 1
}
