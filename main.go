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
			// find a better funciton than .Title
			name: strings.Title(strings.Join(t[1:], " ")),
			year: t[0],
		})
	}

	results := make([][]scrape.Player, 0)

	var wg = sync.WaitGroup{}
	for _, team := range ts {
		wg.Add(1)
		go func(wg *sync.WaitGroup, tm string, yr string) {
			defer wg.Done()
			yearLink := scrape.FindYrBB(yr)
			teamLink := scrape.FindTeamBB(yearLink, tm)
			ps := scrape.FindPlayers(teamLink)
			results = append(results, ps)
		}(&wg, team.name, team.year)
	}
	wg.Wait()

	fmt.Println(results[0][0])
}

type Team struct {
	name string
	year string
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
