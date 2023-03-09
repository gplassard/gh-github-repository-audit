package checks

import (
	"fmt"
	"github.com/cli/go-gh/pkg/repository"
	"github.com/fatih/color"
	"github.com/gplassard/gh-github-repository-audit/utils"
	"os"
)

func CheckRepoVisibility(repo repository.Repository) bool {
	client := utils.RestClient()
	response := struct{ Visibility string }{}
	err := client.Get("repos/"+repo.Owner()+"/"+repo.Name(), &response)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	isValid := response.Visibility == "private"
	if isValid {
		fmt.Println(color.GreenString("Repo has a valid visibility (" + response.Visibility + ") ✅"))
	} else {
		fmt.Println(color.RedString("Repo has invalid visibility (" + response.Visibility + ") ❌"))
	}
	return isValid
}
