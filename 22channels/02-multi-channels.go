package main

import (
	"fmt"
)

func main() {
	// Create three unbuffered string channels
	mychannel := make(chan string)
	anotherChannel := make(chan string)
	yetAnotherChannel := make(chan string)

	// Start a goroutine that sends "data" to mychannel
	go func() {
		mychannel <- "data"
	}()
	// Start a goroutine that sends "cow" to anotherChannel
	go func() {
		anotherChannel <- "cow"
	}()
	// Start a goroutine that sends "goat" to yetAnotherChannel
	go func() {
		yetAnotherChannel <- "goat"
	}()

	// Use select to wait for a message from any of the three channels
	// Only the first message received will be printed
	select {
	case msgFromChannel := <-mychannel:
		fmt.Println(msgFromChannel) // Print message from mychannel
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel) // Print message from anotherChannel
	case msgFromYetAnotherChannel := <-yetAnotherChannel:
		fmt.Println(msgFromYetAnotherChannel) // Print message from yetAnotherChannel
	}

}
