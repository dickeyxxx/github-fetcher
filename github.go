package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/go-github/github"

	"code.google.com/p/goauth2/oauth"
)

var GithubRepoOwner = os.Getenv("GITHUB_REPO_OWNER")
var GithubRepoName = os.Getenv("GITHUB_REPO_NAME")

var githubConfig = &oauth.Config{
	ClientId:     os.Getenv("GITHUB_KEY"),
	ClientSecret: os.Getenv("GITHUB_SECRET"),
	RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
	Scope:        "admin:repo_hook",
	AuthURL:      "https://github.com/login/oauth/authorize",
	TokenURL:     "https://github.com/login/oauth/access_token",
}

func GithubClient(token string) *github.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	return github.NewClient(t.Client())
}

type GithubEvent struct {
	Ref        string            `json:"ref"`
	Before     string            `json:"before"`
	After      string            `json:"after"`
	Repository *GithubRepository `json:"repository"`
}

type GithubRepository struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

func GithubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	event := &GithubEvent{}
	json.Unmarshal(b, &event)
	fmt.Fprintf(w, "%#v", event)
}
