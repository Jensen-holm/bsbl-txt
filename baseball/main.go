package main

import (
	"bufio"
	"fmt"
	"github.com/Jensen-holm/SportSimulation/bbref"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
)

func main() {

	ts := make([]*bbref.Team, 0)

	c := cases.Title(language.AmericanEnglish)

	// create team structs
	for i := 0; i < 2; i++ {
		t := strings.Split(CLInput(), " ")
		name := c.String(strings.Join(t[1:], " "))
		yr := t[0]
		ts = append(ts, bbref.NewTeam(name, yr))
	}

	bbref.GetTeams(ts)

	for _, tm := range ts {
		tm.EstimateRotation()
		tm.EstimateLineup()
		// the estimate lineup situation when trying to consider
		// position is tricky
		// but I think we want users to set lineups
		// or we scrape the lineups before each game anyway
		// link - > https://www.lineups.com/mlb/lineups
	}

	for i := 0; i < 16200; i++ {
		err := bbref.Game(ts[0], ts[1], ts[1].Pitchers()[0], ts[0].Pitchers()[0], 1)
		if err != nil {
			panic(err)
		}
		fmt.Println(ts[0].Wins(), ts[1].Wins())
	}

}

// make a simulation func

func CLInput() string {
	fmt.Println("Enter Team: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.Replace(input, "\n", "", 1)
}
