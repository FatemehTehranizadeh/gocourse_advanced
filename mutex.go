package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(factorialWithGoroutine(6))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func factorialWithGoroutine(n int) int {
	res := 1
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	// wg.Add(n)
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(j int) {
			// time.Sleep(time.Millisecond * 200)
			fmt.Printf("Gorotine %d is starting...\n", j)
			mutex.Lock()
			res *= j
			mutex.Unlock()
			fmt.Printf("The result of Goroutin %d is %d\n", j, res)
			wg.Done()
		}(i)

	}
	wg.Wait()
	// time.Sleep(time.Second)
	return res
}
