package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type LogEntry struct {
	RequestID    string
	Method       string
	Path         string
	Status       int
	ResponseTime float64
	Dyno         string
	ConnectTime  float64
}

type DynoStats struct {
	TotalRequests int
	TotalResponse float64
	MaxConnect    float64
}

func parseLogLine(line string) (*LogEntry, error) {
	// Check if line contains required string and status=200
	if !strings.Contains(line, "coderbyte heroku/router") {
		return nil, fmt.Errorf("not a router log")
	}
	if !strings.Contains(line, "status=200") {
		return nil, fmt.Errorf("status is not 200")
	}

	entry := &LogEntry{}

	// Extract request_id
	if match := regexp.MustCompile(`request_id=(\S+)`).FindStringSubmatch(line); len(match) > 1 {
		entry.RequestID = match[1]
	}

	// Extract method
	if match := regexp.MustCompile(`method=(\S+)`).FindStringSubmatch(line); len(match) > 1 {
		entry.Method = match[1]
	}

	// Extract path
	if match := regexp.MustCompile(`path="([^"]+)"`).FindStringSubmatch(line); len(match) > 1 {
		entry.Path = match[1]
	}

	// Extract status
	if match := regexp.MustCompile(`status=(\d+)`).FindStringSubmatch(line); len(match) > 1 {
		entry.Status, _ = strconv.Atoi(match[1])
	}

	// Extract response time (service=XXXms)
	if match := regexp.MustCompile(`service=(\d+)ms`).FindStringSubmatch(line); len(match) > 1 {
		entry.ResponseTime, _ = strconv.ParseFloat(match[1], 64)
	}

	// Extract dyno
	if match := regexp.MustCompile(`dyno=(\S+)`).FindStringSubmatch(line); len(match) > 1 {
		entry.Dyno = match[1]
	}

	// Extract connect time
	if match := regexp.MustCompile(`connect=(\d+)ms`).FindStringSubmatch(line); len(match) > 1 {
		entry.ConnectTime, _ = strconv.ParseFloat(match[1], 64)
	}

	return entry, nil
}

func processLogs(urlToParse string) error {
	// Step 1: Make GET request
	resp, err := http.Get(urlToParse)
	if err != nil {
		return fmt.Errorf("error fetching URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %v", resp.StatusCode)
	}

	// Step 2: Read response body
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	// Step 3: Split into lines
	lines := strings.Split(string(response), "\n")

	// Step 4: Group by dyno
	dynoMap := make(map[string]*DynoStats)

	// Step 5: Process each line
	for _, line := range lines {
		entry, err := parseLogLine(line)
		if err != nil {
			continue // Skip lines that don't match criteria
		}

		// Initialize dyno stats if not exists
		if _, exists := dynoMap[entry.Dyno]; !exists {
			dynoMap[entry.Dyno] = &DynoStats{}
		}

		// Update stats
		stats := dynoMap[entry.Dyno]
		stats.TotalRequests++
		stats.TotalResponse += entry.ResponseTime

		// Update max connect time
		if entry.ConnectTime > stats.MaxConnect {
			stats.MaxConnect = entry.ConnectTime
		}
	}

	// Step 6: Calculate and print results
	for dyno, stats := range dynoMap {
		avgResponse := stats.TotalResponse / float64(stats.TotalRequests)
		fmt.Printf("%s: %d, %.2f ms, %.2f ms\n",
			dyno,
			stats.TotalRequests,
			avgResponse,
			stats.MaxConnect)
	}

	return nil
}

func main() {
	if err := processLogs("https://coderbyte.com/api/challenges/logs/web-logs-raw"); err != nil {
		log.Fatal("Application failed: ", err)
	}
}
