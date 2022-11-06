package scrape

// could probably change this to bbref, and separate this one and savant

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func HandleGetRequest(url string, r *http.Response, err error) {
	if err != nil {
		panic(err)
	}
	if r.StatusCode != 200 {
		log.Fatalf("odd response status code: %v\n Url: %s", r.StatusCode, url)
	}
}

var bbrefPrefix = "https://baseball-reference.com"

func FindYrBB(year string) string {
	var def = "https://baseball-reference.com/leagues/"
	r, err := http.Get(def)
	HandleGetRequest(def, r, err)

	defer r.Body.Close()
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

func FindTeamBB(yearHref string, team string) string {
	r, err := http.Get(yearHref)
	HandleGetRequest(yearHref, r, err)

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		panic(err)
	}

	var teamHref string
	table := doc.Find("tbody")
	table.Find("th").Each(func(i int, s1 *goquery.Selection) {
		txt := s1.Find("a").Text()
		if txt == team {
			href, ok := s1.Find("a").Attr("href")
			if ok {
				teamHref = bbrefPrefix + href
			}
		}
	})
	return teamHref
}

func FindPlayers(teamHref string) []string {

	r, err := http.Get(teamHref)
	HandleGetRequest(teamHref, r, err)

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		panic(err)
	}

	table := doc.Find("tbody")

	table.Find("td").Each(func(i int, s1 *goquery.Selection) {

		fmt.Println(s1.Text())

	})

	s := make([]string, 0)
	return s

}
