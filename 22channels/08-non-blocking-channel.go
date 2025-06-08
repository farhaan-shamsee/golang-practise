package main

import "fmt"

func main() {
	// Create an unbuffered channel for strings
	messages := make(chan string)
	// Create an unbuffered channel for bools
	signals := make(chan bool)

	// Non-blocking receive: if there's a message, receive it; otherwise, execute default
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received") // No message yet, so this runs
	}

	msg := "hi"
	// Non-blocking send: try to send msg, if not possible, execute default
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent") // No receiver, so this runs
	}

	// Non-blocking multi-channel select: try to receive from messages or signals, else default
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity") // No activity on either channel, so this runs
	}
}

/*
Notes for revision:
- Unbuffered channels block unless both sender and receiver are ready.
- The select statement with a default case makes channel operations non-blocking.
- If no case in select is ready, default is executed immediately.
- Useful for checking channel state without blocking program execution.
*/
