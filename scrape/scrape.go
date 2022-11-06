package scrape

import (
	. "github.com/Jensen-holm/SportSimulation/bsbl"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"sync"
)

// HandleGetRequest -> we want to add headers in the future
func HandleGetRequest(str string, url string, r *http.Response, err error) {
	if err != nil && len(url) == 0 {
		log.Fatalf("URL for '%s' not found: %v", str, url)
	}
	if err != nil {
		panic(err)
	}
	if r.StatusCode != 200 {
		log.Fatalf("odd response status code: %v\n Url: %s", r.StatusCode, url)
	}
}

var bbPrefix = "https://baseball-reference.com"

func FindYrBB(year string) string {
	var def = "https://baseball-reference.com/leagues/"
	r, err := http.Get(def)
	HandleGetRequest(year, def, r, err)

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
				yrHref = bbPrefix + link
			}
		})
	})
	return yrHref
}

func FindTeamBB(yearHref string, team string) string {
	r, err := http.Get(yearHref)
	HandleGetRequest(team, yearHref, r, err)

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
				teamHref = bbPrefix + href
			}
		}
	})
	return teamHref
}

func FindPlayers(teamName string, teamHref string) ([]*Player, []*Player) {
	r, err := http.Get(teamHref)
	HandleGetRequest(teamName, teamHref, r, err)

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		panic(err)
	}

	batTbl := doc.Find("table#team_batting")
	pitTbl := doc.Find("table#team_pitching")
	pitchers := ParseBBTbl(pitTbl)
	hitters := ParseBBTbl(batTbl)
	return pitchers, hitters
}

func ParseBBTbl(tbl *goquery.Selection) []*Player {
	players := make([]*Player, 0)
	cols := make([]string, 0)
	tbl.Each(func(i int, tbl *goquery.Selection) {
		tbl.Find("thead").Each(func(j int, thead *goquery.Selection) {
			thead.Find("tr").Each(func(k int, r *goquery.Selection) {
				r.Find("th").Each(func(l int, th *goquery.Selection) {
					if th.Text() != "Rk" {
						// excluding Rk because it is not found when searching through td tags
						cols = append(cols, th.Text())
					}
				})
			})
		})
		tbl.Find("tbody").Find("tr").Each(func(i int, row *goquery.Selection) {
			p := make(map[string]string, 0)
			row.Find("td").Each(func(j int, td *goquery.Selection) {
				if j < len(cols) {
					p[cols[j]] = td.Text()
				}
			})
			np := new(Player)
			np.SetData(p)
			players = append(players, np)
		})
	})
	return players
}

func GetTeams(teams []*Team) {
	data := make(map[string][]*Player)
	var wg = sync.WaitGroup{}
	for _, team := range teams {
		wg.Add(1)
		go func(wg *sync.WaitGroup, tm *Team) {
			defer wg.Done()
			yearLink := FindYrBB(tm.Year())
			teamLink := FindTeamBB(yearLink, tm.Name())
			hs, ps := FindPlayers(tm.Name(), teamLink)
			data[tm.Name()+" Hitters"] = hs
			data[tm.Name()+" Pitchers"] = ps
		}(&wg, team)
	}
	wg.Wait()
	for k, v := range data {
		for _, tm := range teams {
			if strings.Contains(k, tm.Name()) {
				if strings.Contains(k, "Pitchers") {
					tm.SetPitchers(v)
				} else {
					tm.SetHitters(v)
				}
			}
		}
	}
}
