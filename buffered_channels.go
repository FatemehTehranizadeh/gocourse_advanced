package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 5)
	ls := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for _, v := range ls {
			ch <- v
			fmt.Println("data:", v, "time:", time.Now().Format("15:04:05.000000"))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		for v := range ch {
			fmt.Println(v)
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("The end of main.")

}
