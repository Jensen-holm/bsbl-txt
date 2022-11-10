package main

import (
	"bufio"
	"fmt"
	. "github.com/Jensen-holm/SportSimulation/bbref"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strconv"
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

	for _, tm := range ts {
		tm.EstimateLineup()
		tm.EstimateRotation()
	}

	numSims := SimsInput()

	for i := 0; i < numSims; i++ {

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

func SimsInput() int {
	fmt.Println("\nNumber of Simulations: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0
	}
	if n, err := strconv.Atoi(input); err == nil {
		return n
	}
	return 0
}
