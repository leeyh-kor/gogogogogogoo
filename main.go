package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatalln(err)
	}
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status", res.Status)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	doc.Find(".pagination")
	if err != nil {
		log.Fatalln(err)
	}
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status", res.Status)
	}
	return 0

}
