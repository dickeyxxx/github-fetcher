package main

import (
	"os"

	"github.com/codegangsta/negroni"
)

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
		port = "8080"
	}
	return port
}
