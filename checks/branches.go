package checks

import (
	"fmt"
	"github.com/cli/go-gh/pkg/repository"
	"github.com/fatih/color"
	"github.com/gplassard/gh-github-repository-audit/utils"
	"os"
	"strings"
)

func CheckBranchesPresence(repo repository.Repository) bool {
	isValid := true
	isValid = checkBranchPresence(repo, "dev")
	isValid = checkBranchPresence(repo, "test")
	isValid = checkBranchPresence(repo, "prod")
	return isValid
}

func checkBranchPresence(repo repository.Repository, branch string) bool {
	client := utils.RestClient()
	response := struct{ Name string }{}
	err := client.Get("repos/"+repo.Owner()+"/"+repo.Name()+"/branches/"+branch, &response)

	if err != nil && !strings.HasPrefix(err.Error(), "HTTP 404:") {
		fmt.Println(err)
		os.Exit(1)
	}
	isValid := err == nil
	if isValid {
		fmt.Println(color.GreenString("Repo has a required branch (" + branch + ") ✅"))
	} else {
		fmt.Println(color.RedString("Repo is missing required branch (" + branch + ") ❌"))
	}
	return isValid
}

func CheckBranchesProtection(repo repository.Repository) bool {
	isValid := true
	isValid = checkBranchProtection(repo, "dev")
	isValid = checkBranchProtection(repo, "test")
	isValid = checkBranchProtection(repo, "prod")
	return isValid
}

func checkBranchProtection(repo repository.Repository, branch string) bool {
	client := utils.RestClient()
	response := struct{ Name string }{}
	err := client.Get("repos/"+repo.Owner()+"/"+repo.Name()+"/branches/"+branch+"/protection", &response)

	if err != nil && !strings.HasPrefix(err.Error(), "HTTP 404:") {
		fmt.Println(err)
		os.Exit(1)
	}
	isValid := err == nil
	if isValid {
		fmt.Println(color.GreenString("Repo has required branch protection (" + branch + ") ✅"))
	} else {
		fmt.Println(color.RedString("Repo is missing required branch protection (" + branch + ") ❌"))
	}
	return isValid
}
