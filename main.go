package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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
	mainc := make(chan []extractedJob)
	var jobs []extractedJob

	total := getPages()
	_, bmin, bsec := time.Now().Clock()
	for i := 0; i < total; i++ {
		go getPage(i, mainc)
	}
	for i := 0; i < total; i++ {
		jobs = append(jobs, <-mainc...)
	}
	fmt.Println(jobs)
	_, amin, asec := time.Now().Clock()
	writeJobs(jobs)
	fmt.Println("play_min = " + strconv.Itoa(amin-bmin) + " play_sec = " + strconv.Itoa(asec-bsec))
}

func getPage(page int, mainc chan []extractedJob) {
	c := make(chan extractedJob)
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("wating..." + pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractjob(card, c)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainc <- jobs
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

func extractjob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find("salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}
	w.Write(headers)
	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
	wErr := w.Write(headers)
	checkErr(wErr)

}
