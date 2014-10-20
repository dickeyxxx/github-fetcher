package main

import (
	"log"
	"os"

	"github.com/codegangsta/negroni"
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
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseHandler(Router())
	n.Run(":" + port())
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
