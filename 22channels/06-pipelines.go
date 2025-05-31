package main

import "fmt"

// ***IMPORTANT***
// As we are using unbuffered channel the sliceTOchannel send 1st data and waits for someone(sq channel) to pick it up. and this goes on. 
// The sq channel also waits for new value in the channel.
// Once the loop is complete on sliceTOchannel it send out close signal which triggers the sq chanel to close itself.

// sliceToChannel takes a slice of ints and returns a channel that emits each element.
func sliceToChannel(data []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, r := range data {
			out <- r // Send each element to the channel
		}
		close(out) // Close the channel when done
	}()
	return out
}

// sq takes a channel of ints and returns a channel that emits the square of each int.
func sq(sqChan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for r := range sqChan {
			out <- r * r // Square each received value and send to output channel
		}
		close(out) // Close the channel when done
	}()
	return out
}

func main() {
	nums := []int{1, 2}
	// stage 1: Convert slice to channel
	dataChannel := sliceToChannel(nums)
	// stage 2: Square each value from the previous channel
	finalChannel := sq(dataChannel)

	// stage 3: Print each squared value from the final channel
	for data := range finalChannel {
		fmt.Println(data)
	}
}
