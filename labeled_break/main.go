package main

func main() {
line4:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(i, ", ", j, " \n")
			break line4
		}
	}

	println()

}
