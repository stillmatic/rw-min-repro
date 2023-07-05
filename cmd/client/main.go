package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// pull env var down, hit their ping endpoint every minute
	c := http.Client{}
	targetHost := os.Getenv("TARGET_HOST")

	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				res, err := c.Get(targetHost + "/ping")
				if err != nil {
					log.Println("ERROR: ", err)
				}
				resp, err := io.ReadAll(res.Body)
				if err != nil {
					log.Println("ERROR: ", err)
				}
				log.Println("INFO: ", string(resp))
			}
		}
	}()

}
