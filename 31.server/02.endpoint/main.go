/*
"Write a GET /clusters endpoint that accepts query parameters:
- region (optional)
- status (optional)

Filter results based on provided parameters.

Example:
GET /clusters?region=eu-west-1&status=active
→ Returns clusters matching both filters

*/

package main

import (
	"encoding/json"
	"net/http"
)

type Clusters struct {
	ID     string
	Region string
	Status string
}

func getClusters(w http.ResponseWriter, r *http.Request) {

	region := r.URL.Query().Get("region")
	status := r.URL.Query().Get("status")

	allClusters := []Clusters{
		{ID: "c1", Region: "eu-west-1", Status: "active"},
		{ID: "c2", Region: "eu-west-2", Status: "active"},
		{ID: "c3", Region: "ap-south-1", Status: "inactive"},
	}

	matchedClusters := []Clusters{}

	for _, value := range allClusters {
		if (region == "" || value.Region == region) && (status == "" || value.Status == status) {
			matchedClusters = append(matchedClusters, value)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matchedClusters)
}

func main() {

	http.HandleFunc("/clusters", getClusters)
	http.ListenAndServe(":8080", nil)
}
