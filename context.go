package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
/* Generator Pattern: 
Creating required channels
Creating a goroutine with an anonymous function
returning the channels
*/

func main() {
	//context.Background is used only in main function
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	ch := make(chan string, 7)
	wg := sync.WaitGroup{}
	wg.Add(1)

	write := func(ctx context.Context) {
		defer wg.Done()
		defer close(ch)
		ls := []string{"Zan", "Zendegi", "Azadi"}
		time.Sleep(time.Second*4)
		for _, v := range ls {
			select {
			case <-ctx.Done():
				return
			case ch <- v:
				fmt.Println(v)
			}
		}
	}

	read := func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println(v)
			}
		}
	}

	// go write(ctxWithTimeout)
	go func() {
		// time.Sleep(time.Second * 3)
		write(ctxWithTimeout)
	}()
	// go read()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read(ctxWithTimeout)
	}

	wg.Wait()
	fmt.Println("The end.")
}
