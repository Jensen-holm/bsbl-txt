package savant

import (
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

	tms := make(map[string]string, 0)
	doc.Find("div .team-selector").Each(func(i int, selection *goquery.Selection) {
		selection.Find("select .team-nav").Each(func(i int, selection *goquery.Selection) {
			if id, ok := selection.Attr("id"); ok && id == "ddlteam" {
				tm := selection.Find("option").Text()
				val, ok := selection.Attr("value")
				if ok {
					tms[tm] = val
				}
			}
		})
	})
	return tms, nil
}

func FindTeam(tm string, yr string) (string, error) {
	m, err := TmIdMap()
	if err != nil {
		return "", err
	}
	tmLink := startUrl[:len(startUrl)-3] + m[tm] + "?season=" + yr
	return tmLink, nil
}
