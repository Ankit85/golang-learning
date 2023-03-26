package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Learning channel")

	// declaring channel
	myChan := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// Reading channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChan
		fmt.Printf("value %v \n", val)
		fmt.Printf("isChannelOpen %v \n", isChannelOpen)

		wg.Done()
	}(myChan, wg)
	// sending to channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChan <- 5
		close(myChan)
		wg.Done()
	}(myChan, wg)

	wg.Wait()
}
