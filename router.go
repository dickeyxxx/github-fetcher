package main

import "github.com/gorilla/mux"

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/github/webhook", GithubWebhookHandler)
	return router
}
