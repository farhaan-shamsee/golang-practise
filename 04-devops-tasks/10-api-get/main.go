package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Responses struct {
	Name        string    `json:"name"`
	Reaction    Reactions `json:"reactions,omitempty"`
	PublishedAt time.Time `json:"published_at"`
}

type Reactions struct {
	Hooray int `json:"hooray,omitempty"`
}

// func fetchResponses(url string) ([]Responses, error) { // without context timeout
func fetchResponses(ctx context.Context, url string) ([]Responses, error) {
	client := &http.Client{Timeout: 10 * time.Second} //we use pointer bcz we want to modify the client object, and we want to avoid copying the entire struct.

	// req, err := http.NewRequest(http.MethodGet, url, nil) // without context timeout
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("building request %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("calling API %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var responses []Responses
	if err := json.NewDecoder(resp.Body).Decode(&responses); err != nil {
		return nil, fmt.Errorf("decoding responses: %w", err)
	}

	return responses, nil
}

func fetchAllResponses(ctx context.Context, baseURL string) ([]Responses, error) {
	var allResponses []Responses
	page := 1

	for {
		url := fmt.Sprintf("%s?page=%d&per_page=100", baseURL, page)
		batch, err := fetchResponses(ctx, url)
		if err != nil {
			return nil, err
		}

		if len(batch) == 0 {
			break
		}

		allResponses = append(allResponses, batch...)
		page++
	}

	return allResponses, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure the context is canceled to avoid resource leaks
	responses, err := fetchResponses(ctx, "https://api.github.com/repos/kubernetes/kubernetes/issues/123456/comments")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	last7days := time.Now().AddDate(0, 0, -7)
	responseCount := 0
	for _, response := range responses {
		if response.PublishedAt.After(last7days) {
			responseCount += 1
		}
	}
	fmt.Printf("Number of responses in the last 7 days: %d\n", responseCount)
}
