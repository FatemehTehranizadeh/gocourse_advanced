package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "Zakariya"
		ch2 <- "Tajik"
	}()

	go func() {
		ch2 <- "Khiyal"
		time.Sleep(time.Second * 2)
		ch1 <- "Setareh"
	}()

MAHSA:
	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		case <- time.After(time.Second * 3):
			fmt.Println("I don't wait anymore!")
			break MAHSA
		}
	}

}
