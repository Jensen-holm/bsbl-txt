package bbref

import (
	"net/http"
	"strings"
	"sync"

	"github.com/Jensen-holm/SportSimulation/scrape"
	"github.com/PuerkitoBio/goquery"
)

var (
	bbPrefix = "https://baseball-reference.com"
	def      = bbPrefix + "/leagues/"
)

func FindYrBB(year string) (string, error) {
	r, err := http.Get(def)
	scrape.HandleGetRequest(year, def, r, err)

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return "", err
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
	return yrHref, nil
}

func FindTeamBB(yearHref string, team string) (string, error) {
	r, err := http.Get(yearHref)
	scrape.HandleGetRequest(team, yearHref, r, err)

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return "", err
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
	return teamHref, nil
}

func FindPlayers(teamName string, teamHref string) ([]*Player, []*Player, error) {
	r, err := http.Get(teamHref)
	scrape.HandleGetRequest(teamName, teamHref, r, err)

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, nil, err
	}

	p := ParseBBTbl(doc.Find("table#team_pitching"))
	h := ParseBBTbl(doc.Find("table#team_batting"))

	hitters := make([]*Player, 0)
	for _, hitter := range h {
		hitters = append(hitters, NewPlayer(hitter))
	}

	pitchers := make([]*Player, 0)
	for _, pitcher := range p {
		pitchers = append(pitchers, NewPlayer(pitcher))
	}

	return pitchers, hitters, nil
}

func ParseBBTbl(tbl *goquery.Selection) []map[string]string {
	players := make([]map[string]string, 0)
	cols := make([]string, 0)
	tbl.Each(func(i int, table *goquery.Selection) {
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
				p[cols[j]] = td.Text()
			})
			players = append(players, p)
		})
	})
	return players
}

// GetTeams -> Concurrently calls the functions above to scrape baseball reference
func GetTeams(teams []*Team) {

	var wg = sync.WaitGroup{}
	// need to make sure that we need a results channel
	// to have each routine check to make sure that we don't
	// do the same iteration more than once
	for _, team := range teams {
		wg.Add(1)
		go func(wg *sync.WaitGroup, tm *Team) {
			defer wg.Done()

			yearLink, err := FindYrBB(tm.Year)
			if err != nil {
				panic(err)
			}

			teamLink, err := FindTeamBB(yearLink, tm.Name)
			if err != nil {
				panic(err)
			}

			ps, hs, err := FindPlayers(tm.Name, teamLink)
			if err != nil {
				panic(err)
			}

			tm.SetHitters(hs)
			tm.SetPitchers(ps)

		}(&wg, team)
	}
	wg.Wait()
}
