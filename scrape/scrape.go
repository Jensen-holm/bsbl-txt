package scrape

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"sync"
)

func FindTeamBB(wg *sync.WaitGroup, team string, year string) string {
	defer wg.Done()

	var def = "https://baseball-reference.com/leagues/"
	r, err := http.Get(def)
	if err != nil {
		log.Fatalf("Error Getting baseball reference url: %v", err)
	}

	if r.StatusCode != 200 {
		log.Fatalf("Failed to retrieve data: status code %v", r.StatusCode)
	}

	defer r.Body.Close() // don't think this would ever run into an error
	doc, err := goquery.NewDocumentFromReader(r.Body)

	// find where the link text is equal to the year string
	table := doc.Find("tbody")

	return table.Text()
}
