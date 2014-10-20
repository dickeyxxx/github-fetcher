package main

import (
	"net/http"
	"path"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views/index.html"))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		panic(err)
	}
}
