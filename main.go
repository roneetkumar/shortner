package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	pathsToURL := map[string]string{
		"/urlshort-firestore": "https://godoc.org/cloud.google.com/go/firestore",
	}

	mapHandler := MapHandler(pathsToURL, mux)

	jString := `
	[
		{
			"path": "/url-fire",
			"url": "https://godoc.org/cloud.google.com/go/firestore"
		},
		{
			"path": "/url-my",
			"url": "https://roneetkumar.github.io"
		}
	]`

	jsonHandler, err := JSONHandler(jString, mapHandler)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Starting the server on : 8080")

	// JSON HANDLER
	http.ListenAndServe(":8080", jsonHandler)

	// MAP HANDLER
	// http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
}
