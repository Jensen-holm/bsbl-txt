package main

import (
	"bufio"
	"fmt"
	. "github.com/Jensen-holm/SportSimulation/bsbl"
	. "github.com/Jensen-holm/SportSimulation/scrape"
	"math/rand"
	"os"
	"strings"
	"sync"
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
	r1 := make([]string, 0)
	st1 := time.Now()

	for i := 0; i < 10000; i++ {
		r, _ := PA(ts[0].Hitters()[0], ts[1].Pitchers()[rand.Intn(len(ts[1].Pitchers()))])
		r1 = append(r1, r)
	}

	dur1 := time.Since(st1)
	fmt.Println(dur1)

	st2 := time.Now()

	var wg sync.WaitGroup
	wg.Add(10000)
	results := make([]string, 0)

	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			r, _ := PA(ts[0].Hitters()[0], ts[1].Pitchers()[0])
			fmt.Println(r)
			results = append(results, r)
		}()
	}
	wg.Wait()

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
