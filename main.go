package main

import (
	"bufio"
	"fmt"
	. "github.com/Jensen-holm/SportSimulation/bsbl"
	. "github.com/Jensen-holm/SportSimulation/scrape"
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

	data := make(map[string][]*Player)

	var wg = sync.WaitGroup{}
	for _, team := range ts {
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
		for _, tm := range ts {
			if strings.Contains(k, tm.Name()) {
				if strings.Contains(k, "Pitchers") {
					tm.SetPitchers(v)
				} else {
					tm.SetHitters(v)
				}
			}
		}
	}
	fmt.Println(ts[0])
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
