package main

import (
	"fmt"
	"time"
)

func main() {
	// Print a welcome message
	fmt.Println("welcome to time study of golang")

	// Get the current local time
	presentTime := time.Now()
	// Print the current time
	fmt.Println("present time is", presentTime)

	// Format and print the current time in a custom layout
	// Layout: "01-02-2006 15:04:05 Monday" (MM-DD-YYYY HH:MM:SS Day)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Create a specific date and time: 10th October 2023, 23:00:00 UTC
	createdDate := time.Date(2023, time.October, 10, 23, 0, 0, 0, time.UTC)
	// Print the created date
	fmt.Println("created date is", createdDate)
}
