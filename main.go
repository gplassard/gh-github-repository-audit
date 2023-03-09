package main

import (
	"fmt"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
	"github.com/cli/go-gh/pkg/repository"
	"github.com/fatih/color"
	"os"
)

func main() {
	repo, err := gh.CurrentRepository()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Starting audit for repo", color.YellowString(repo.Owner()+"/"+repo.Name()))

	valid := true
	valid = valid && checkTeamTopic(err, repo)

	if valid {
		fmt.Println(color.GreenString("All checks have passed, congratulations ! ✅"))
	} else {
		fmt.Println(color.RedString("Some checks have failed ❌"))
		os.Exit(1)
	}
}

func checkTeamTopic(err error, repo repository.Repository) bool {
	opts := api.ClientOptions{
		Headers: map[string]string{"Accept": "application/vnd.github+json", "X-GitHub-Api-Version": "2022-11-28"},
	}
	client, err := gh.RESTClient(&opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	response := struct{ Names []string }{}
	err = client.Get("repos/"+repo.Owner()+"/"+repo.Name()+"/topics", &response)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	isValid := hasValidTeamTopic(response.Names)
	if isValid {
		fmt.Println(color.GreenString("Repo has a valid team topic ✅"))
	} else {
		fmt.Println(color.RedString("Repo has no valid team topic ❌"))
	}
	return isValid
}

func hasValidTeamTopic(topics []string) bool {
	acceptedTags := make(map[string]bool)
	acceptedTags["team-gplassard"] = true
	for _, topic := range topics {
		if acceptedTags[topic] {
			return true
		}
	}
	return false
}
