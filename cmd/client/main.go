package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/rs/xid"
)

type response struct {
	Body string `json:"body"`
}

func main() {
	// pull env var down, hit their ping endpoint every minute
	c := http.Client{}
	targetHost := os.Getenv("TARGET_HOST")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		guid := xid.New()
		target := fmt.Sprintf("%s/ping", targetHost)
		url, err := url.Parse(target)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req := &http.Request{
			Method: "GET",
			URL:    url,
			Header: http.Header{
				"X-Request-ID": []string{guid.String()},
			},
		}
		res, err := c.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()
		// decode to json
		var resp response
		err = json.NewDecoder(res.Body).Decode(&resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		returnedID := res.Header.Get("X-Request-ID")
		log.Println(fmt.Sprintf("[INFO] Got response to %s: %s", returnedID, resp.Body))
		fmt.Fprintf(w, resp.Body)

	})
	// this is public so needs a port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	port = fmt.Sprintf(":%s", port)
	fmt.Printf("Starting client on port %s, hitting %s\n", port, targetHost)
	http.ListenAndServe(port, nil)
}
