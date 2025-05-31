// package main // Declare the main package

// import (
// 	"fmt" // Import the fmt package for formatted I/O
// 	"time"
// )

// func main() { // Main function

// 	// This is an infinite channel which will run till the time is up.

// 	go func ()  {
// 		for {
// 			select {
// 			default:
// 				fmt.Println("DOING WORK")
// 			}
// 		}
		
// 	}()

// 	time.Sleep(time.Second * 3)

// }


package main

import (
    "fmt"
    "time"
)

func worker(jobs <-chan int, id int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second)
    }
}

func main() {
    jobs := make(chan int)
    for i := 1; i <= 3; i++ {
        go worker(jobs, i)
    }
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    time.Sleep(time.Second * 3)
}
