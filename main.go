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
	title := os.Getenv("INPUT_TEXT")
	desp := os.Getenv("INPUT_DESP")
	openid := os.Getenv("INPUT_OPENID")

	if key == "" || title == "" {
		log.Fatalln("key or title[text] is required")
	}

	if len(title) > 32 {
		log.Fatalln("title[text] is longer than 32 characters")
	}

	reqURL := fmt.Sprintf("https://sctapi.ftqq.com/%s.send", key)
	res, err := http.PostForm(reqURL, url.Values{
		"title": []string{title},
		"desp": []string{desp},
		"openid": []string{openid},
	})
	if err != nil {
		log.Fatalln("post error:", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalln("status code:", res.StatusCode)
	}
}
