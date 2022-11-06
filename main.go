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
			name: strings.Join(t[1:], " "),
			year: t[0],
		})
	}

	p := make([][]string, 0)

	var wg = sync.WaitGroup{}
	for _, team := range ts {
		wg.Add(1)
		go func(wg *sync.WaitGroup, tm string, yr string) {
			defer wg.Done()
			yearLink := scrape.FindYrBB(yr)
			teamLink := scrape.FindTeamBB(yearLink, tm)
			playerLinks := scrape.FindPlayers(teamLink)

			// does go automatically flatten lists ??
			p = append(p, playerLinks)

		}(&wg, team.name, team.year)
	}
	wg.Wait()

	// flatten it
	fp := Flatten(p)
	fmt.Println(fp)

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

func Flatten(sl [][]string) []string {
	f := make([]string, 0)
	for _, i := range sl {
		for _, j := range i {
			f = append(f, j)
		}
	}
	return f
}
