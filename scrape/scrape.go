package scrape

import (
	"log"
	"net/http"
	"strconv"
	"unicode"
)

import (
	. "github.com/Jensen-holm/SportSimulation/bsbl"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"sync"
)

func IsLetter(s string) bool {
	// this line is pretty ghetto
	s = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "*", ""), ".", ""), "#", ""), "-", "")
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

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

func FindYrBB(year string) (string, error) {
	var def = "https://baseball-reference.com/leagues/"
	r, err := http.Get(def)
	HandleGetRequest(year, def, r, err)

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
	HandleGetRequest(team, yearHref, r, err)

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

func FindPlayers(teamName string, teamHref string) (map[string]map[string]string, map[string]map[string]string, error) {
	r, err := http.Get(teamHref)
	HandleGetRequest(teamName, teamHref, r, err)

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, nil, err
	}

	batTbl := doc.Find("table#team_batting")
	pitTbl := doc.Find("table#team_pitching")
	pitchers := ParseBBTbl(pitTbl)
	hitters := ParseBBTbl(batTbl)
	return pitchers, hitters, nil
}

func ParseBBTbl(tbl *goquery.Selection) []*Player {
	players := make([]*Player, 0)
	ps := make(map[string]map[string]string, 0)
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
				p[cols[j]] = td.Text()
			})
			ps[p["Name"]] = p
		})
	})

	// create player objects
	for p, _ := range ps {

		np := Player{
			PA:  strconv.ParseInt(ps[p]["PA"], 1, 64),
			AB:  0,
			R:   0,
			H:   0,
			B2:  0,
			B3:  0,
			HR:  0,
			RBI: 0,
			SB:  0,
			CS:  0,
			BB:  0,
			SO:  0,
			TB:  0,
			GBP: 0,
			HBP: 0,
			SH:  0,
			SF:  0,
			IBB: 0,
			W:   0,
			L:   0,
			WP:  0,
			BK:  0,
			BF:  0,
			SV:  0,
			ER:  0,
			IP:  0,
		}
		players = append(players, &np)
	}
	return players
}

func GetTeams(teams []*Team) {
	data := make([]*Player, 0)
	var wg = sync.WaitGroup{}

	for _, team := range teams {
		wg.Add(1)
		go func(wg *sync.WaitGroup, tm *Team) {
			defer wg.Done()

			yearLink, err := FindYrBB(tm.Year())
			if err != nil {
				panic(err)
			}

			teamLink, err := FindTeamBB(yearLink, tm.Name())
			if err != nil {
				panic(err)
			}

			hs, ps, err := FindPlayers(tm.Name(), teamLink)
			if err != nil {
				panic(err)
			}

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
