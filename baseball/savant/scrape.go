package savant

import (
	"fmt"
	"github.com/Jensen-holm/SportSimulation/scrape"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

var startUrl = "https://baseballsavant.mlb.com/team/113"

func TmIdMap() (map[string]string, error) {
	r, err := http.Get(startUrl)
	scrape.HandleGetRequest("Reds", startUrl, r, err)

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return map[string]string{}, err
	}

	tms := doc.Find("select#ddlteam")
	fmt.Println(tms.Text())
	return map[string]string{}, nil
}

func FindTeam(tm string, yr string) (string, error) {
	m, err := TmIdMap()
	if err != nil {
		return "", err
	}
	tmLink := startUrl[:len(startUrl)-3] + m[tm] + "?season=" + yr
	return tmLink, nil
}
