// Package main provides ...
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/build", build)
	http.HandleFunc("/push", push)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
