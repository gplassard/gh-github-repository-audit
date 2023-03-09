package main

import (
	"fmt"
	"github.com/cli/go-gh"
	"github.com/fatih/color"
	"github.com/gplassard/gh-github-repository-audit/checks"
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
	valid = checks.CheckTeamTopic(repo) && valid
	valid = checks.CheckRepoVisibility(repo) && valid

	if valid {
		fmt.Println(color.GreenString("All checks have passed, congratulations ! ✅"))
	} else {
		fmt.Println(color.RedString("Some checks have failed ❌"))
		os.Exit(1)
	}
}
