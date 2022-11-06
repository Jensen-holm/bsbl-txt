package scrape

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func HandleGetRequest(url string, r *http.Response, err error) {
	if err != nil {
		log.Fatalf("error requesting '%s': %v", url, err)
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
	doc.Find("tbody").Find("tr").Each(func(i int, s1 *goquery.Selection) {
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
	doc.Find("tbody").Find("th").Each(func(i int, s1 *goquery.Selection) {
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

func FindPlayers(teamHref string) []map[string]string {
	r, err := http.Get(teamHref)
	HandleGetRequest(teamHref, r, err)

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		panic(err)
	}

	tbls := make([]map[string]string, 0)

	tables := doc.Find("table")

	// Columns
	cols := make([]string, 0)
	tables.Each(func(i int, tbl *goquery.Selection) {
		tbl.Find("thead").Each(func(j int, thead *goquery.Selection) {
			thead.Find("tr").Each(func(k int, r *goquery.Selection) {
				r.Find("th").Each(func(l int, th *goquery.Selection) {
					cols = append(cols, th.Text())
				})
			})
		})
	})

	// data
	tables.Each(func(i int, tbl *goquery.Selection) {
		tbl.Find("tbody").Find("tr").Each(func(j int, r *goquery.Selection) {
			d := make(map[string]string, 0)
			r.Find("td").Each(func(k int, td *goquery.Selection) {
				d[cols[i]] = td.Text()
				tbls = append(tbls, d)
			})
		})
	})
	return tbls
}
