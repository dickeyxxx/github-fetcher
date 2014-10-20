package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
	git "github.com/libgit2/git2go"
)

var GithubToken = os.Getenv("GITHUB_TOKEN")
var GithubRepoOwner = os.Getenv("GITHUB_REPO_OWNER")
var GithubRepoName = os.Getenv("GITHUB_REPO_NAME")
var RepoPath = path.Join("tmp", GithubRepoOwner, GithubRepoName)
var GithubUrl = "git://github.com/" + GithubRepoOwner + "/" + GithubRepoName

func GithubClient(token string) *github.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	return github.NewClient(t.Client())
}

func listCommits() {
	client := GithubClient(GithubToken)
	commits, _, err := client.Repositories.ListCommits(GithubRepoOwner, GithubRepoName, nil)
	if err != nil {
		panic(err)
	}
	for _, commit := range commits {
		fmt.Println(*commit.Commit.Committer.Date)
	}
}

func main() {
	_, err := os.Stat(RepoPath)
	if os.IsNotExist(err) {
		cloneRepo()
	}
	os.Chdir(RepoPath)
	getLatest()
	runMake()
	fmt.Println("done")
}

func getLatest() {
	repo, err := git.OpenRepository(".")
	if err != nil {
		panic(err)
	}
	remote, err := repo.LoadRemote("origin")
	if err != nil {
		panic(err)
	}
	err = remote.Fetch(nil, nil, "")
	if err != nil {
		panic(err)
	}
}

func runMake() {
	fmt.Println("Building...")
	cmd := exec.Command("make")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func cloneRepo() {
	fmt.Printf("Cloning %s...\n", GithubUrl)
	err := os.MkdirAll(path.Dir(RepoPath), 0755)
	if err != nil {
		panic(err)
	}
	cloneOptions := &git.CloneOptions{}
	_, err = git.Clone(GithubUrl, RepoPath, cloneOptions)
	if err != nil {
		panic(err)
	}
}
