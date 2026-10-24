package main

import (
	"fmt"
	"net/url"
)

func urlMaker(baseUrl string, clusterId string, event string, metadata map[string]string) (string, error) {
	newUrl, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	newUrl.Path = "/webhook"
	params := url.Values{}
	params.Add("cluster_id", clusterId)
	params.Add("event", event)

	for key, value := range metadata {
		params.Add(key, value)
	}

	newUrl.RawQuery = params.Encode()

	return newUrl.String(), nil

}

func main() {
	metadata := map[string]string{
		"region": "eu-west-1",
		"owner":  "team-x",
	}

	webhookURL, _ := urlMaker(
		"https://hooks.scout24.com",
		"prod-eu-123",
		"cluster.created",
		metadata,
	)

	fmt.Println(webhookURL)
}
