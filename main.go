// Package main provides ...
package main

import (
	"log"
	"net/http"

	"github.com/alacine/deptool/handlers"
)

func main() {
	http.HandleFunc("/status", handlers.Status)
	http.HandleFunc("/upload", handlers.Upload)
	http.HandleFunc("/build", handlers.Build)
	http.HandleFunc("/push", handlers.Push)
	http.HandleFunc("/clean", handlers.Clean)
	http.HandleFunc("/dict", handlers.Dict)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
