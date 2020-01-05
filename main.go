package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	key := os.Getenv("INPUT_KEY")
	text := os.Getenv("INPUT_TEXT")
	desp := os.Getenv("INPUT_DESP")

	if key == "" || text == "" {
		log.Fatalln("key or text is required")
	}

	reqMsg := &url.Values{
		"text": []string{text},
		"desp": []string{desp},
	}

	reqURL := fmt.Sprintf("https://sc.ftqq.com/%s.send", key)
	res, err := http.Post(reqURL, "application/x-www-form-urlencoded", strings.NewReader(reqMsg.Encode()))
	if err != nil {
		log.Fatalln("post error:", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalln("status code:", res.StatusCode)
	}
}
