package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("present time is", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //This is the format for time in golang, we always have to use this format

	createdDate := time.Date(2023, time.October, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("created date is", createdDate)
}
