package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Jensen-holm/SportSimulation/bbref"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		fmt.Println(tm.Hitters()[0])
		//	// the estimate lineup situation when trying to consider
		//	// position is tricky
		//	// but I think we want users to set lineups
		//	// or we scrape the lineups before each game anyway
		//	// link - > https://www.lineups.com/mlb/lineups
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
