package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Responses struct {
    Name string `json:"name"`
    Reaction Reactions `json:"reactions,omitempty"`
    PublishedAt time.Time `json:"published_at"`
}

type Reactions struct {
    Hooray int `json:"hooray,omitempty"`
}

func main() {
    var response []Responses
    resp, err := http.Get("https://api.github.com/repos/backstage/backstage/releases")
    if err != nil {
        fmt.Println("error: %s", err)
    }
    if err:= json.NewDecoder(resp.Body).Decode(&response); err != nil {
        fmt.Println("error: %s", err)
    }
    resp.Body.Close()
    twoMonthsAgo := time.Now().AddDate(0,-2,0)
    hoorayCount := 0
    for _, formattedResp := range response {
        if formattedResp.PublishedAt.After(twoMonthsAgo) {
            hoorayCount = hoorayCount+formattedResp.Reaction.Hooray
        }

    }
    fmt.Println(hoorayCount)
}