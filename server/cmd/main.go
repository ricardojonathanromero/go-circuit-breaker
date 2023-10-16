package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/ping", ping)
	fmt.Println("Server Application running.")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println("error server", err)
	}
}

func ping(w http.ResponseWriter, _ *http.Request) {
	fprintf, err := fmt.Fprintf(w, "pong from server-go")
	if err != nil {
		fmt.Println("error response", fprintf)
	}
}
