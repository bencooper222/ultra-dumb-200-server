package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read PORT environment variable and validate it's a port number
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT env var must be set")
	}

	urlPathStr := os.Getenv("URL_PATH")
	if urlPathStr == "" {
		urlPathStr = "/"
		log.Println("URL_PATH env var not set, using default '/'")
	}
	if !strings.HasPrefix(urlPathStr, "/") {
		log.Fatal("URL_PATH env var must start with '/'")
	}
	// we just convert to an int to check it is a number
	// we use portStr below since that's what the http package wnats
	_, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("PORT env var must be a number")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == urlPathStr {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	err = http.ListenAndServe(":"+portStr, nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed")
	} else if err != nil {
		log.Fatal("Couldn't start server on port " + portStr)
	}
}
