package main

import (
	"fmt"
	"gitlab-go"
	"os"
)

func main() {
	gitlabAddress := os.Args[1]
	privateToken := os.Args[2]
	client := gitlab.Connect(gitlabAddress, privateToken)
	project, err := client.GetProject("")
	if err != nil {
		panic(err)
	}
	fmt.Println(project)
}
