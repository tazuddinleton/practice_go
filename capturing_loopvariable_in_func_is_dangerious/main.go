package main

import "sync"

func main() {

}

func problem2() {
	const N = 10
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)

	for i := 0; i < N; i++ {
		// loop variables captured by go func may have
		// unexpected values.
		go func(idx int) {
			defer wg.Done()
			mu.Lock()
			m[idx] = idx // Using parameter is safest
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	println(len(m))
}
