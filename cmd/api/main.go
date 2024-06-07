package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bekk/github-actions-workshop/internal/greeting"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	names := queryValues["name"]
	test := queryValues["test"]

	if len(names) > 0 {
		greeting, err := greeting.Greet(names)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Internal error\n")
		} else {
			io.WriteString(w, greeting)
		}
	} else {
		io.WriteString(w, "Hello, no one?\n")
		io.WriteString(w, fmt.Sprintf("test: %s\n", test))
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got request to %s\n", r.URL.Path)
	io.WriteString(w, "Hello world\n")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)

	port := "8888"
	fmt.Printf("Running on port %s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
