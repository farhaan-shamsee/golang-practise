package main

import (
	"fmt"
	"log"

	"gitlab.com/gitlab-org/api/client-go"
)

const gitlabToken string = "my-token"

func main() {
	fmt.Println("This is gitlab client")

	git, err := gitlab.NewClient(gitlabToken, gitlab.WithBaseURL("https://gitlab.com/api/v4"))
	if err != nil {
		log.Fatalf("Unable to create client: ", err)
	}

	// git.Projects.ListProjects()
	// fmt.Println(git.Projects.ListProjects(&gitlab.ListProjectsOptions{}))

	projectId := 459259
	trueVal := true
	deleteProject, err := git.Projects.DeleteProject(projectId, &gitlab.DeleteProjectOptions{PermanentlyRemove: &trueVal})
	if err != nil {
		log.Fatalf("Unable to create client: ", err)
	}
	fmt.Printf("status: %v", deleteProject.Response.Status)
	fmt.Printf("header: %v", deleteProject.Response.Header)

}
