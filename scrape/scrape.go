package scrape

import (
	"log"
	"net/http"
)

func FindTeamBB(team string, year string) *http.Response {
	var def = "https://baseball-reference.com/leagues/"
	r, err := http.Get(def)
	if err != nil {
		log.Fatalf("Error Getting baseball reference url: %v", err)
	}
	return r
}
