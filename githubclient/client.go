package githubclient

/*
Module that houses the client. Saves 1 instance which is accessible across the CLI
*/

import (
	"os"

	"github.com/google/go-github/v62/github"
)

var (
	token  = os.Getenv("GITHUB_TOKEN")
	client *github.Client
)

// Initalise the client with the necessary token. TODO: Export token to .env variable
func InitClient() {
	if client == nil {
		client = github.NewClient(nil).WithAuthToken(token)
	}
}

// Function to get the client
func GetClient() *github.Client {
	return client
}

func GetToken() string {
	return token
}
