package main

import (
	"bufio"
	"fmt"
	"github.com/Jensen-holm/SportSimulation/bbref"
	"log"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {

	teams, err := GetInput()
	if err != nil {
		log.Fatalf("error getting team input: %v", err)
	}

	bbref.GetTeams(teams)
	TeamSetUp(teams)

}

// CLInput -> prompts for and scans user input
// for baseball teams to simulate against each other
func CLInput() string {
	fmt.Println("Enter Team: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.Replace(input, "\n", "", 1)
}

// GetInput -> Takes raw user input from CLI and creates
// a baseball reference team object out of it. need to add
// a check to see if they are real teams
func GetInput() ([]*bbref.Team, error) {
	var c = cases.Title(language.AmericanEnglish)

	tms := make([]*bbref.Team, 0)
	for i := 0; i < 2; i++ {
		t := strings.Split(CLInput(), " ")
		name := c.String(strings.Join(t[1:], " "))
		yr := t[0]
		tms = append(tms, bbref.NewTeam(name, yr))
	}
	return tms, nil
}

// TeamSetUp -> Runs functions necessary to set the
// lineups and rotations for each team we scraped
func TeamSetUp(tms []*bbref.Team) {
	for _, tm := range tms {
		tm.EstimateRotation()
		tm.EstimateLineup()
	}
}
