package main

import (
	"bufio"
	"fmt"
	. "github.com/Jensen-holm/SportSimulation/bsbl"
	. "github.com/Jensen-holm/SportSimulation/scrape"
	"math/rand"
	"os"
	"strings"
	"time"
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

	// simulate lots of games
	st1 := time.Now()
	for i := 0; i < 100000; i++ {
		PA(ts[0].Hitters()[0], ts[1].Pitchers()[rand.Intn(len(ts[1].Pitchers()))])
	}

	dur1 := time.Since(st1)
	fmt.Printf("Without Go Routines: %v\n", dur1)

	st2 := time.Now()

	for i := 0; i < 100000; i++ {
		go PA(ts[0].Hitters()[0], ts[1].Pitchers()[0])
	}

	dur2 := time.Since(st2)
	fmt.Printf("With go routines: %v", dur2)

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
