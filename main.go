package main

import (
	"bufio"
	"fmt"
	"github.com/Jensen-holm/SportSimulation/bbref"
	"github.com/fatih/color"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	teams, err := Teams("Enter Team name -> ")
	if err != nil {
		log.Fatalf("error getting team input: %v", err)
	}

	bbref.GetTeams(teams)
	TeamSetUp(teams)

	sims, err := NumSims("Enter number of simulations -> ")
	if err != nil {
		log.Fatalf("error getting number of simulations: %v", err)
	}

	teams, err = bbref.Simulation(sims, teams)
	if err != nil {
		log.Fatalf("error in bbref simulation function -> %v", err)
	}

	Results(teams, sims)

}

// CLInput -> prompts for and scans user input
// for baseball teams to simulate against each other
// kind of not cool that we ignore an error in this
// but shouldn't run into it
func CLInput(prompt string) string {
	color.Cyan("\n" + prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.Replace(input, "\n", "", 1)
}

// Teams -> Takes raw user input from CLI and creates
// a baseball reference team object out of it. need to add
// a check to see if they are real teams
func Teams(prompt string) ([]*bbref.Team, error) {
	var c = cases.Title(language.AmericanEnglish)

	tms := make([]*bbref.Team, 0)
	for i := 0; i < 2; i++ {
		t := strings.Split(CLInput(prompt), " ")
		name := c.String(strings.Join(t[1:], " "))
		yr := t[0]
		tms = append(tms, bbref.NewTeam(name, yr))
	}
	return tms, nil
}

// NumSims -> Sole responsibility is getting the number of
// simulations to perform from the user
func NumSims(prompt string) (int64, error) {
	num := CLInput(prompt)
	if n, err := strconv.ParseInt(num, 0, 64); err != nil {
		return 0, fmt.Errorf(
			"could not convert '%s' into an integer: %v", num, err,
		)
	} else {
		return n, nil
	}
}

// TeamSetUp -> Runs functions necessary to set the
// lineups and rotations for each team we scraped
func TeamSetUp(tms []*bbref.Team) {
	for _, tm := range tms {
		tm.EstimateRotation()
		tm.EstimateLineup()
	}
}

func Results(teams []*bbref.Team, sims int64) {
	for _, team := range teams {
		color.Green(
			"\n%s %s win percentage: %.2f\n",
			team.Year(),
			team.Name(),
			float64(team.Wins())/float64(sims)*100,
		)
	}
}
