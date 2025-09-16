package main

import (
	"fmt"
)

func main() {
	receiveOnlyCh := generator()
	for v := range receiveOnlyCh {
		fmt.Println(v)
	}

}

func generator() <-chan string {
	ch := make(chan string)
	go func() {
		ls := []string{"Mahsa Amini", "Sarina Esmaeilzadeh", "Abolfazl Adinehzadeh", "Abolfazl Mehdipour", "Matin Naderkhani", "Nika Shakarami"}
		for _, v := range ls {
			ch <- v
		}
		defer close(ch)
	}()

	return ch
}
