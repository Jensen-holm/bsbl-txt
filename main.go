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
	return input
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

	r := make(chan string)

	var wg = sync.WaitGroup{}

	for _, team := range ts {
		wg.Add(1)
		go func(wg *sync.WaitGroup, team string, year string) {
			table := scrape.FindTeamBB(wg, team, year)
			r <- table
		}(&wg, team.name, team.year)
	}

	wg.Wait()
	fmt.Println(r)
}
