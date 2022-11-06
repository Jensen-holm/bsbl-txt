package main

import (
	"bufio"
	"fmt"
	"github.com/Jensen-holm/SportSimulation/scrape"
	"os"
	"strings"
)

func CLInput() string {
	fmt.Println("Enter Team: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return input
}

type Team struct {
	name string
	year string
}

func main() {

	ts := make([]Team, 0)

	for i := 0; i < 2; i++ {
		t := strings.Split(CLInput(), " ")
		ts = append(ts, Team{
			name: strings.Join(t[1:], " "),
			year: t[0],
		})
	}

	table := scrape.FindTeamBB(" ", " ")
	fmt.Println(table)

	r := make(chan string)

	for _, t := range ts {
		go func() {
			resp := scrape.FindTeamBB(t.name, t.year)
			fmt.Println(&resp)
			r <- resp
		}()
	}

}
