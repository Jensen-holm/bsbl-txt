package main

import (
	"github.com/Jensen-holm/SportSimulation/bbref"
	"github.com/marcusolsson/tui-go"
	"log"
	"strconv"
)

func main() {

	team1 := tui.NewEntry()
	team2 := tui.NewEntry()
	numSims := tui.NewEntry()

	form := tui.NewGrid(0, 0)
	form.AppendRow(
		tui.NewLabel("Team 1"),
		tui.NewLabel("Team 2"),
		tui.NewLabel("Num Sims"),
	)
	form.AppendRow(team1, team2, numSims)

	status := tui.NewStatusBar("Ready.")

	init := tui.NewButton("[Go]")
	init.OnActivated(func(b *tui.Button) {
		GoButton(team1, team2, numSims, status)
	})

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, init),
	)

	window := tui.NewVBox(
		tui.NewPadder(12, 0, tui.NewLabel("Welcome to bsbl.txt! Enter teams and number of simulations to get started")),
		tui.NewPadder(1, 1, form),
		buttons,
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	root := tui.NewVBox(
		content,
		status,
	)

	tui.DefaultFocusChain.Set(team1, team2, numSims, init)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err = ui.Run(); err != nil {
		log.Fatal(err)
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

// GoButton -> Runs when the Go button is pressed and
// valid teams are entered as input
func GoButton(team1, team2, sims *tui.Entry, status *tui.StatusBar) {

	status.SetText("Simulating ...")
	nSims, _ := strconv.ParseInt(sims.Text(), 0, 64)

	// create teams
	status.SetText("Scraping team data ...")
	tms := []*bbref.Team{
		bbref.NewTeam(team1.Text()),
		bbref.NewTeam(team2.Text()),
	}

	// collect team data
	bbref.GetTeams(tms)

	// start simulation with the teams
	TeamSetUp(tms)
	tms, err := bbref.Simulation(nSims, tms)
	if err != nil {
		log.Fatalf("error in the simulation: %v", err)
	}

	status.SetText("Done!")
}
