package scrape

import (
	"fmt"
	"log"
	"net/http"
)

func FindTeamBB(team string, year string) {
	var def = "https://baseball-reference.com/leagues/"
	html, err := http.Get(def)
	if err != nil {
		log.Fatalf("Error Getting baseball reference url: %v", err)
	}
	fmt.Println(html)
}
