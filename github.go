package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func GithubWebhookHandler(res http.ResponseWriter, req *http.Request) {
	b, _ := ioutil.ReadAll(req.Body)
	event := &GithubEvent{}
	json.Unmarshal(b, &event)
	fmt.Fprintf(res, "%#v", event)
}
