package utils

import (
	"fmt"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
	"os"
)

func RestClient() api.RESTClient {
	opts := api.ClientOptions{
		Headers: map[string]string{"Accept": "application/vnd.github+json", "X-GitHub-Api-Version": "2022-11-28"},
	}
	client, err := gh.RESTClient(&opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return client
}
