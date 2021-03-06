package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	c := make(chan bool)
	total := getPages()
	fmt.Println(time.Now().Clock())
	for i := 0; i < total; i++ {
		go getPage(i, c)
		fmt.Println(<-c)
	}
	fmt.Println(time.Now().Clock())
}

func getPage(page int, c chan bool) chan bool {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting pageURL: ", pageURL)
	c <- true
	return c
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	pages := 0
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	if err != nil {
		log.Fatalln(err)
	}
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status", res.Status)
	}
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkCode(code *http.Response) {
	if code.StatusCode != 200 {
		log.Fatalln("Request failed with Status", code.Status)
	}
}
