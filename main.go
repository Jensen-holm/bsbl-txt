package main

import (
	"bufio"
	"fmt"
	. "github.com/Jensen-holm/SportSimulation/bsbl"
	. "github.com/Jensen-holm/SportSimulation/scrape"
	"os"
	"strings"
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
	GetTeams(ts)

	ts[0].Hitters()[0].PaProbsHit()

	games := make([]*Game, 10000)
	for i := 0; i < 16200; i++ {
		ng := new(Game)
		ng.SetHome(ts[0])
		ng.SetAway(ts[1])
		games = append(games, ng)
	}

	// play a ton of games at once with go routines

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
