package checks

import (
	"fmt"
	"github.com/cli/go-gh/pkg/repository"
	"github.com/fatih/color"
	"github.com/gplassard/gh-github-repository-audit/utils"
	"os"
)

func CheckRepoPermissions(repo repository.Repository) bool {
	client := utils.RestClient()
	response := []struct {
		Name       string
		Permission string
	}{}
	err := client.Get("repos/"+repo.Owner()+"/"+repo.Name()+"/teams", &response)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	isValid := true
	isValid = checkRepoTeam(response, "dev", "admin")
	isValid = checkRepoTeam(response, "operations", "admin")
	return isValid
}

func checkRepoTeam(teams []struct {
	Name       string
	Permission string
}, expectedTeam string, expectedPermission string) bool {
	for _, team := range teams {
		if team.Name == expectedTeam {
			if team.Permission == expectedPermission {
				fmt.Println(color.GreenString("Team " + team.Name + " has valid permission (" + team.Permission + ") ✅"))
				return true
			} else {
				fmt.Println(color.RedString("Team " + team.Name + " has invalid permission (" + team.Permission + ") ❌"))
				return false
			}
		}
	}
	fmt.Println(color.RedString("Repo is missing permission for team " + expectedTeam + " ❌"))
	return false
}
