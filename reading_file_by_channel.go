package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {

	ch := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	file, err := os.Open("/home/reera/gocourse/intermediate/exampleFile.txt")
	if err != nil {
		log.Fatal("Error while opening the app:", err)
	}
	fmt.Println("File is opened successfully!")

	var buff bytes.Buffer
	scanner := bufio.NewScanner(file)

	go func() {
		defer wg.Done()
		for scanner.Scan() {
			line := scanner.Text()

			ch <- line
		}
	}()

	go func() {
		wg.Done()
		var c int
		for v := range ch {
			_, err = io.WriteString(&buff, v)
			if err != nil {
				log.Fatal("Error while writing to the buffer:", err)
			}
			c++
			fmt.Printf("The content of the line %d is %s\n",c, v)
		}
	}()

	wg.Wait()
	fmt.Println("The buffer is:",buff.String())
	defer file.Close()

}
