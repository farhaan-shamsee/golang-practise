package main

import (
	"fmt"
	"sync"
	// "log"
	"net/http"
	// "time"
)

var signals = []string{"test"}

var wg sync.WaitGroup //usually these are pointers
var mut sync.Mutex //usually these are pointers

func main() {
	// go greeter("hello") // we made this a go routine
	// greeter("world")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() //always comes at the end of the main func, so that it waits for the whole flow to complete
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		//if we do not add this then the thread where we sent the go routine does not come back and we wont get the output
// 		// not an ideal solution
// 		time.Sleep(3 * time.Millisecond)

// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("OOPS in endpoint: %s\n", endpoint)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}


}
