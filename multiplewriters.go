package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		ls := []string{"Mahsa Amini", "Sarina Esmaeilzadeh", "Abolfazl Adinehzadeh", "Abolfazl Mehdipour", "Matin Naderkhani", "Nika Shakarami"}
		for _, v := range ls {
			ch <- v
		}
		wg.Done()
	}()

	go func() {
		ls := []string{"Pedram Azarnoush", "Ayda Rostami", "Amir Djawadifar"}
		for _, v := range ls {
			ch <- v
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

}
