package main

func main() {
	problem14()
}
func problem1() {
	// guess the output
	s := "123"
	ps := &s
	b := []byte(*ps)
	pb := &b

	s += "4"
	*ps += "5"
	b[1] = '0'

	print(*ps)
	print(string(*pb))

}

func problem14() {
	// guess the output
	const N = 3
	m := make(map[int]*int)
	for i := 0; i < N; i++ {
		m[i] = &i
	}

	for _, v := range m {
		println(*v)
	}
}
