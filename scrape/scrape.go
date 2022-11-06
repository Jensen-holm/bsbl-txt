package scrape

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"time"
)

func HandleGetRequest(url string, r *http.Response, err error) {
	if err != nil {
		log.Fatalf("error requesting '%s': %v", url, err)
	}
	if r.StatusCode != 200 {
		log.Fatalf("odd response status code: %v\n Url: %s", r.StatusCode, url)
	}
}

func ResponseToTable(r *http.Response) *goquery.Selection {
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatalf("error creating new goquery document: %v", err)
	}
	return doc.Find("tbody")
}

var bbrefPrefix = "https://baseball-reference.com"

func FindYrBB(year string) string {
	var def = "https://baseball-reference.com/leagues/"
	time.Sleep(1)
	r, err := http.Get(def)
	HandleGetRequest(def, r, err)

	defer r.Body.Close()
	doc := ResponseToTable(r)

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
	doc := ResponseToTable(r)

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

func FindPlayers(teamHref string) []map[string]string {

	r, err := http.Get(teamHref)
	HandleGetRequest(teamHref, r, err)

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		panic(err)
	}

	tbls := make([]map[string]string, 0)

	doc.Find("tbody").Each(func(i int, tbl *goquery.Selection) {

		cols := make([]string, 0)
		tbl.Find("thead").Find("th").Each(func(j int, col *goquery.Selection) {
			cols = append(cols, col.Text())
		})

		dm := make(map[string]string)
		tbl.Find("td").Each(func(k int, td *goquery.Selection) {
			dm[cols[k]] = td.Text()
		})
		tbls[i] = dm
	})
	return tbls
}
