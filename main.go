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

	c := cases.Title(language.AmericanEnglish)

	// create team structs
	for i := 0; i < 2; i++ {
		t := strings.Split(CLInput(), " ")
		nt := new(Team)
		nt.SetName(c.String(strings.Join(t[1:], " ")))
		nt.SetYear(t[0])
		ts = append(ts, nt)
	}

	// scrape player data with go routines
	// and assign them with their corresponding teams
	GetTeams(ts)

	for _, tm := range ts {
		tm.EstimateLineup()
		tm.EstimateRotation()
	}

	// 16,200 sims by default
	var ab int
	for i := 0; i < 16200; i++ {
		r, j, err := HalfInning(ab, ts[0], ts[0].Pitchers()[0])
		if err != nil {
			panic(err)
		}
		ab = j
		fmt.Println(r)
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
