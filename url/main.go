package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{url, len(body)}
}

func main() {
	pages := make(chan Page)
	urls := []string{"https://golang.org/", "https://golang.org/doc", "https://example.com"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for range len(urls) {
		fmt.Println(<-pages)
	}
}
