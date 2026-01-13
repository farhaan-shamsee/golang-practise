// This is an infinite channel which will run till the time is up.
/*
package main

import (
	"fmt"
	"time"
)

func main() {
	go func ()  {
		for {
			select {
			default:
				fmt.Println("DOING WORK")
			}
		}
	}()
	time.Sleep(time.Second * 3)
}
*/

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

// This program demonstrates a worker pool pattern using Go channels and goroutines.

// What it does:
// Creates a channel called jobs to send integer jobs to workers.
// Starts 3 worker goroutines (worker(jobs, i)) that listen on the jobs channel.
// Sends 5 jobs (integers 1 to 5) into the jobs channel.
// Closes the channel after sending all jobs, signaling workers to stop when no more jobs are available.
// Waits 3 seconds (time.Sleep(time.Second * 3)) to allow workers to finish processing before the program exits.
// Output Example
// Worker 1 processing job 1
// Worker 2 processing job 2
// Worker 3 processing job 3
// Worker 1 processing job 4
// Worker 2 processing job 5
// Summary:
// The program distributes 5 jobs among 3 workers, each worker processes jobs as they arrive, demonstrating concurrent job processing with channels and goroutines.
