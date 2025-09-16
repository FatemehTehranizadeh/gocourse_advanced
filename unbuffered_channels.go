package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)

	write := func() {
		defer wg.Done()
		ls := []string{"Zan", "Zendegi", "Azadi", "Mahsa", "Siavash", "Nika", "Matin"}
		for _, v := range ls {
			// time.Sleep(time.Second * 2)
			ch <- v
			// fmt.Println(time.Now())
		}
		defer func() {
			close(ch)
			fmt.Println("Channel is closed.")
			v, ok := <-ch
			fmt.Println("Reading from channel after it's closed: ", v)
			fmt.Println("Is channel still open? ", ok)
		}()
	}

	read := func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("data:", v, "time:", time.Now().Format("15:04:05.000000"))
			time.Sleep(time.Second * 5)
		}
		fmt.Println("Reader ends its function")
		// time.Sleep(time.Second)
	}

	go write()
	go func() {
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go read()
		}
	}()

	// time.Sleep(time.Second * 1)
	wg.Wait()
	fmt.Println("The end.", time.Now().Format("15:04:05.000000"))
}
