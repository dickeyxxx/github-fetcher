package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

var GithubRepoOwner = os.Getenv("GITHUB_REPO_OWNER")
var GithubRepoName = os.Getenv("GITHUB_REPO_NAME")
var RepoPath = path.Join("tmp", GithubRepoOwner, GithubRepoName)
var GithubUrl = "git://github.com/" + GithubRepoOwner + "/" + GithubRepoName

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
	runCmd("git", "fetch", "origin", "master")
	runCmd("git", "merge", "origin/master", "--ff")
}

func runMake() {
	runCmd("make")
}

func runCmd(name string, args ...string) {
	fmt.Println("Building...")
	cmd := exec.Command(name, args...)
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
	runCmd("git", "clone", GithubUrl, RepoPath)
}
