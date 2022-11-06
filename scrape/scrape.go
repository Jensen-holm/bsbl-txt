package scrape

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func FindTeamBB(team string, year string) string {
	var def = "https://baseball-reference.com/leagues/"
	r, err := http.Get(def)
	if err != nil {
		log.Fatalf("Error Getting baseball reference url: %v", err)
	}

	defer r.Body.Close() // don't think this would ever run into an error
	doc, err := goquery.NewDocumentFromReader(r.Body)

	// find where the link text is equal to the year string
	table := doc.Find("tbody")

	return table.Text()
}
