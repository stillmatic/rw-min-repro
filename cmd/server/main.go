package server

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	msg := os.Getenv("MESSAGE")
	// sets up a server that listens on port 8080
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong -- "+msg)
	})

	http.ListenAndServe(":8080", nil)
}
