package checks

import (
	"fmt"
	"github.com/cli/go-gh/pkg/repository"
	"github.com/fatih/color"
	"github.com/gplassard/gh-github-repository-audit/utils"
	"os"
)

func CheckTeamTopic(repo repository.Repository) bool {
	client := utils.RestClient()
	response := struct{ Names []string }{}
	err := client.Get("repos/"+repo.Owner()+"/"+repo.Name()+"/topics", &response)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hasValidTeamTopic, topic := getTeamTopic(response.Names)
	if hasValidTeamTopic {
		fmt.Println(color.GreenString("Repo has a valid team topic (" + topic + ") ✅"))
	} else {
		fmt.Println(color.RedString("Repo has no valid team topic ❌"))
	}
	return hasValidTeamTopic
}

func getTeamTopic(topics []string) (bool, string) {
	acceptedTags := make(map[string]bool)
	acceptedTags["team-gplassard"] = true
	for _, topic := range topics {
		if acceptedTags[topic] {
			return true, topic
		}
	}
	return false, ""
}
