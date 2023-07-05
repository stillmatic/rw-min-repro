package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type response struct {
	Body string `json:"body"`
}

func main() {
	msg := os.Getenv("MESSAGE")
	fmt.Printf("Starting server with secret message %s\n", msg)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		fmt.Printf("Received request id %s \n", reqID)
		resp := response{Body: msg}
		// return json
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Request-ID", reqID)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe(":8080", nil)
}
