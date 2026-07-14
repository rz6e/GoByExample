package main

import (
	"fmt"
	"sync"
)

func writeInts(messages chan<- int) {

	for i := range 5 {
		messages <- i
	}

	close(messages)

}

func readInts(messages <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := range messages {
		fmt.Println(i)
	}

}

func main() {
	messages := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)

	go writeInts(messages)
	go readInts(messages, &wg)

	wg.Wait()
}
