package main

import (
	"bufio"
	"fmt"
	"github.com/Jensen-holm/SportSimulation/scrape"
	"os"
	"strings"
	"sync"
)

func CLInput() string {
	fmt.Println("Enter Team: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return input[:len(input)-1]
}

type Team struct {
	name string
	year string
}

func main() {

	ts := make([]Team, 0)

	for i := 0; i < 2; i++ {
		t := strings.Split(CLInput(), " ")
		ts = append(ts, Team{
			name: strings.Join(t[1:], " "),
			year: t[0],
		})
	}

	teamData := make([]string, 2)

	var wg = sync.WaitGroup{}
	fmt.Println(len(ts[0].year))

	for _, team := range ts {
		wg.Add(1)
		go func(wg *sync.WaitGroup, team string, year string) {
			defer wg.Done()
			yearLink := scrape.FindYrBB(year)
			teamData = append(teamData, yearLink)
		}(&wg, team.name, team.year)
	}
	wg.Wait()
	fmt.Println(teamData)
}
