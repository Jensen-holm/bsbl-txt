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
		tm.EstimateLineup()
		tm.EstimateRotation()
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
