package scrape

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

var bbrefPrefix = "https://baseball-reference.com"

func FindYrBB(year string) string {
	var def = "https://baseball-reference.com/leagues/"
	r, err := http.Get(def)
	if err != nil {
		log.Fatalf("Error Getting baseball reference url: %v", err)
	}
	if r.StatusCode != 200 {
		log.Fatalf("Failed to retrieve data: status code %v", r.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		panic(err)
	}
	var yrHref string
	table := doc.Find("tbody")
	table.Find("tr").Each(func(i int, s1 *goquery.Selection) {
		s1.Find("th").Each(func(i int, s2 *goquery.Selection) {
			link, ok := s2.Find("a").Attr("href")
			if ok && strings.Contains(link, year) {
				yrHref = bbrefPrefix + link
			}
		})
	})
	return yrHref
}
