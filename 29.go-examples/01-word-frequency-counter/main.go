package main

import (
    "fmt"
    "regexp"
    "strings"
)

func CountWordFrequency(text string) map[string]int {
    // Convert to lowercase for case-insensitive comparison
    text = strings.ToLower(text)
    
    // Extract words: sequences of letters and digits only
    // \w+ matches [a-zA-Z0-9_]+ but we'll filter underscores later
    re := regexp.MustCompile(`\w+`)
    words := re.FindAllString(text, -1)
    
    // Count frequency using map
    frequency := make(map[string]int)
    for _, word := range words {
        // Remove underscores (not considered letters/digits for words)
        cleanWord := strings.Trim(word, "_")
        if cleanWord != "" {
            frequency[cleanWord]++
        }
    }
    
    return frequency
}

func main() {
    // Test cases
    fmt.Println("Test 1:")
    result1 := CountWordFrequency("The quick brown fox jumps over the lazy dog.")
    fmt.Println(result1)
    
    fmt.Println("\nTest 2:")
    result2 := CountWordFrequency("Hello, hello! How are you doing today? Today is a great day.")
    fmt.Println(result2)
}
