package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Buffered channels - send & receive demo")

	// 1️⃣ Buffered channel that can hold 2 ints without blocking.
	myCh := make(chan int, 2)

	// 2️⃣ Wait for sender + receiver.
	var wg sync.WaitGroup
	wg.Add(2)

	//------------------------------------------------
	// 🚀 RECEIVER: range-loops until channel closes
	//------------------------------------------------
	go func(ch <-chan int) {
		defer wg.Done()

		for v := range ch { // stops automatically when ch is closed
			fmt.Printf("  📥 received: %d\n", v)
		}
		fmt.Println("  ✅ channel closed, receiver done")
	}(myCh)

	//------------------------------------------------
	// 🚀 SENDER: pushes some data, then closes channel
	//------------------------------------------------
	go func(ch chan<- int) {
		defer wg.Done()

		toSend := []int{1, 2, 3, 4} // values we’ll transmit
		for _, n := range toSend {
			fmt.Printf("📤 send: %d\n", n)
			ch <- n // blocks only when buffer is full
		}

		close(ch) // signal “no more data”
		fmt.Println("✅ sender closed channel")
	}(myCh)

	//------------------------------------------------
	// ⏳ wait for both goroutines
	//------------------------------------------------
	wg.Wait()
	fmt.Println("🎉 all done")
}
