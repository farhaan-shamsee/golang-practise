package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myCh := make(chan int, 2)

	// the below will give error called "all goroutines are asleep - deadlock!", meaning that unless someone is listening to the
	// channel we can not assign value to the channel. Kind of chicken egg problem
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	// Receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		close(myCh)
		val, isChannelOpen := <- myCh
		
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// Send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(myCh)	// if placed here it gives: "panic: send on closed channel"
		// we can listen to closed channel. It will give 0. But how do we guarantee it is coming from closed channel or myCh <- 0.
		// myCh <- 0
		// myCh <- 5
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
