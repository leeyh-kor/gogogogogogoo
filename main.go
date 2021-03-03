package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	fmt.Println("wating for message")
	for _, url := range urls {
		go hitURL(url, c)

	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is cool"
}

func hitURL(url string, c chan result) {
	// fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "Ok"
	if err != nil || resp.StatusCode >= 400 {
		status = "Fail"
	}
	c <- result{url: url, status: status}
}
