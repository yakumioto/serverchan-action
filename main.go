package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	key := os.Getenv("INPUT_KEY")
	title := os.Getenv("INPUT_TITLE")
	desp := os.Getenv("INPUT_DESP")

	if key == "" || title == "" {
		log.Fatalln("key or title is required")
	}

	reqURL := fmt.Sprintf("https://sctapi.ftqq.com/%s.send", key)
	res, err := http.PostForm(reqURL, url.Values{"title": {title}, "desp": {desp}})
	
	if err != nil {
		log.Fatalln("post error:", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalln("status code:", res.StatusCode)
	}
}
