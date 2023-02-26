package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	scrapeUrl := "https://merolagani.com/CompanyDetail.aspx?symbol=NMBHF1"
	c := colly.NewCollector(colly.AllowedDomains("merolagani.com"))
	c.OnHTML("div.panel-body tbody.panel", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US;q=0")
		fmt.Println(fmt.Sprintf("Visiting %s\n", r.URL))
	})

	c.OnHTML("div.panel-body tbody.panel", func(h *colly.HTMLElement) {
		selection := h.DOM
		fmt.Println(selection.Find("table"))
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while Scraping %s\n", e.Error())
	})

	c.Visit(scrapeUrl)
}
