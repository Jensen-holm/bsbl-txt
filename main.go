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

	// create team structs
	for i := 0; i < 2; i++ {
		t := strings.Split(CLInput(), " ")
		nt := new(Team)
		nt.SetName(strings.Title(strings.Join(t[1:], " ")))
		nt.SetYear(t[0])
		ts = append(ts, nt)
	}

	// scrape player data with go routines
	// and assign them with their corresponding teams
	GetTeams(ts)

	// testing the plate appearance function
	tstHitter := ts[0].Hitters()[10]
	tstPitcher := ts[1].Pitchers()[10]
	fmt.Println(PA(tstHitter, tstPitcher))

	// simulate lots of games

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
