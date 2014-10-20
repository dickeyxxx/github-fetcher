package main

import (
	"fmt"
	"log"
)

func init() {
	if GithubRepoOwner == "" || GithubRepoName == "" {
		log.Fatalln("GITHUB_REPO_OWNER and/or GITHUB_REPO_NAME are not set")
	}
	if githubConfig.ClientId == "" || githubConfig.ClientSecret == "" || githubConfig.RedirectURL == "" {
		log.Fatalln("GITHUB_KEY GITHUB_SECRET, and/or GITHUB_REDIRECT_URL are not set")
	}
}

func main() {
	fmt.Println("yay")
}
