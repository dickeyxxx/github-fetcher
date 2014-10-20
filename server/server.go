package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-github/github"
)

func Run() {
	http.HandleFunc("/github/webhook", GithubHook)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func GithubHook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var push github.PushEvent
	json.Unmarshal(body, &push)
	if err != nil {
		panic(err)
	}
	fmt.Printf("doc %+v\n", *push.Ref)
	fmt.Fprintf(w, "yay")
}
