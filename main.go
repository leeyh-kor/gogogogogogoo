package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	location string
	title    string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	c := make(chan bool)
	total := getPages()
	_, bmin, bsec := time.Now().Clock()
	for i := 0; i < total; i++ {
		go getPage(i, c)
		fmt.Println(<-c)
	}
	_, amin, asec := time.Now().Clock()
	fmt.Println("play_min = " + strconv.Itoa(amin-bmin) + "play_sec = " + strconv.Itoa(asec-bsec))
}

func getPage(page int, c chan bool) chan bool {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("data-jk")
		title := CleanString(card.Find(".title>a").Text())
		location := CleanString(card.Find(".sjcl").Text())
		fmt.Println(id, title, location)
	})

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

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
