package main

import (
	"bufio"
	"fmt"
	"github.com/Jensen-holm/SportSimulation/scrape"
	"os"
	"strings"
	"sync"
)

func main() {

	ts := make([]Team, 0)

	for i := 0; i < 2; i++ {
		t := strings.Split(CLInput(), " ")
		ts = append(ts, Team{
			// find a better function than .Title
			name: strings.Title(strings.Join(t[1:], " ")),
			year: t[0],
		})
	}

	var wg = sync.WaitGroup{}
	for _, team := range ts {
		wg.Add(1)
		go func(wg *sync.WaitGroup, tm Team) {
			defer wg.Done()
			yearLink := scrape.FindYrBB(tm.year)
			teamLink := scrape.FindTeamBB(yearLink, tm.name)
			hs, ps := scrape.FindPlayers(teamLink)
			team.pitchers = ps
			team.hitters = hs
		}(&wg, team)
	}
	wg.Wait()
}

type Team struct {
	name     string
	year     string
	hitters  []scrape.Player
	pitchers []scrape.Player
}

func (tm *Team) SetHitters(hitters []scrape.Player) {
	tm.hitters = hitters
}

func (tm *Team) SetPitchers(pitchers []scrape.Player) {
	tm.pitchers = pitchers
}

func CLInput() string {
	fmt.Println("Enter Team: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.Replace(input, "\n", "", 1)
}
