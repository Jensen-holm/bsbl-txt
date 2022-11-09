package main

import (
	"bufio"
	"fmt"
	. "github.com/Jensen-holm/SportSimulation/bbref"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
)

func main() {

	ts := make([]*Team, 0)

	caser := cases.Title(language.AmericanEnglish)

	// create team structs
	for i := 0; i < 2; i++ {
		t := strings.Split(CLInput(), " ")
		nt := new(Team)
		nt.SetName(caser.String(strings.Join(t[1:], " ")))
		nt.SetYear(t[0])
		ts = append(ts, nt)
	}

	// scrape player data with go routines
	// and assign them with their corresponding teams
	GetTeams(ts)

	ts[0].EstimateLineup()

	for _, p := range ts[0].Lineup() {
		fmt.Println(p.Name())
	}
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
