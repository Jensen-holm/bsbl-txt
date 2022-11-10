package scrape

import (
	"log"
	"net/http"
)

// HandleGetRequest -> we want to add headers in the future
func HandleGetRequest(str string, url string, r *http.Response, err error) {
	if err != nil && len(url) == 0 {
		log.Fatalf("URL related to '%s' not found: %v", str, url)
	}
	if err != nil {
		log.Fatalf("Error requesting url (check internet connection?): Error : %v", err)
	}
	if r.StatusCode != 200 {
		log.Fatalf("odd response status code: %v\n Url: %s", r.StatusCode, url)
	}
}
