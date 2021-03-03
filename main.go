package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	// urls := []string{
	// 	"https://www.airbnb.com/",
	// 	"https://www.google.com/",
	// 	"https://www.amazon.com/",
	// 	"https://www.reddit.com/",
	// 	"https://www.google.com/",
	// 	"https://soundcloud.com/",
	// 	"https://www.facebook.com/",
	// 	"https://www.instagram.com/",
	// 	"https://academy.nomadcoders.co/",
	// }

	c := make(chan string)
	people := [2]string{"nico", "yuhyun"}

	fmt.Println("wating for message")
	for _, person := range people {
		go isSexy(person, c)

	}
	for i := 0; i < len(people); i++ {
		fmt.Println("received message: ", <-c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is cool"
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
