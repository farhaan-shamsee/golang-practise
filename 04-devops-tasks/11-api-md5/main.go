package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func normalizeURL(raw string) string {
	if strings.HasPrefix(raw, "http://") || strings.HasPrefix(raw, "https://") {
		return raw
	}
	return "http://" + raw
}

func fetchAndHash(url string, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- struct{}{}
	defer func() { <-sem }()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading %s: %v\n", url, err)
		return
	}

	hash := md5.Sum(body)
	fmt.Printf("%s %x\n", url, hash)
}

func main() {
	parallel := flag.Int("parallel", 10, "maximum number of parallel requests")
	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("usage: myhttp [--parallel N] url1 url2 ...")
		os.Exit(1)
	}

	sem := make(chan struct{}, *parallel)
	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)
		go fetchAndHash(normalizeURL(u), sem, &wg)
	}

	wg.Wait()
}
