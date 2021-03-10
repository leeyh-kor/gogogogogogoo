package main

import (
	"fmt"
	"gogoscrapper"
	goscrapper "goscraper"
	"scrapper"
	"time"
)

func main() {
	p := fmt.Println
	start := time.Now().Second()
	scrapper.Scrape("python")
	time.Sleep(5)
	vanila_scrapper := time.Now().Second()
	goscrapper.Scrape("c#")
	time.Sleep(5)
	go_scrapper := time.Now().Second()
	gogoscrapper.Scrape("C++")
	time.Sleep(5)
	gogo_scrapper := time.Now().Second()

	p("vanila_runtime", vanila_scrapper-start)
	p("go_runtime", go_scrapper-vanila_scrapper)
	p("gogo_runtime", gogo_scrapper-go_scrapper)

}
