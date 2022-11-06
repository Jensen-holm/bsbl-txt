package main

import (
	"bufio"
	"fmt"
	. "github.com/Jensen-holm/SportSimulation/bsbl"
	"github.com/Jensen-holm/SportSimulation/scrape"
	"os"
	"strings"
	"sync"
)

func main() {

	ts := make([]*Team, 0)

	for i := 0; i < 2; i++ {
		t := strings.Split(CLInput(), " ")
		nt := new(Team)
		nt.SetName(strings.Title(strings.Join(t[1:], " ")))
		nt.SetYear(t[0])
		ts = append(ts, nt)
	}

	var wg = sync.WaitGroup{}
	for _, team := range ts {
		wg.Add(1)
		go func(wg *sync.WaitGroup, tm *Team) {
			defer wg.Done()
			yearLink := scrape.FindYrBB(tm.GetYear())
			teamLink := scrape.FindTeamBB(yearLink, tm.GetName())
			hs, ps := scrape.FindPlayers(tm.GetName(), teamLink)
			team.SetPitchers(ps) // not sure what this warning is about
			team.SetHitters(hs)
		}(&wg, team)
	}
	wg.Wait()
	fmt.Println(ts)
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
