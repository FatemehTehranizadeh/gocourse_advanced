package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan string, 7)
	wg := sync.WaitGroup{}
	wg.Add(1)

	write := func() {
		defer wg.Done()
		ls := []string{"Zan", "Zendegi", "Azadi", "Mahsa", "Siavash", "Nika", "Matin"}
		for _, v := range ls {
			ch <- v
		}
		defer close(ch)

	}

	read := func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}

	go write()
	// go read()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	fmt.Println("The end.")
}
