package main

import "time"

func main() {
	//goroutines
	// two or more threads running at once
	// can access the same variable at the same time
	// this can cause unexpected behavior
	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	m2 := make(map[int]int)

	go func() {
		for i := 0; i < 1000; i++ {
			m2[i] = i
		}
	}()

	go func() {
		for i := 1001; i < 2000; i++ {
			m2[i] = i
		}
	}()

	time.Sleep(5 * time.Second)

}
