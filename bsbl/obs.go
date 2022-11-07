package bsbl

// Player -> Contains Simple Data about a player scraped from baseball reference
type Player struct {
	data map[string]string
}

func (p *Player) SetData(d map[string]string) {
	p.data = d
}

// Team -> Contains Data about a team
type Team struct {
	name     string
	year     string
	hitters  []*Player
	pitchers []*Player
}

func (tm *Team) SetName(n string) {
	tm.name = n
}

func (tm *Team) SetYear(yr string) {
	tm.year = yr
}

func (tm *Team) SetHitters(hitters []*Player) {
	tm.hitters = hitters
}

func (tm *Team) SetPitchers(pitchers []*Player) {
	tm.pitchers = pitchers
}

func (tm *Team) Year() string {
	return tm.year
}

func (tm *Team) Name() string {
	return tm.name
}

func (tm *Team) Hitters() []*Player {
	return tm.hitters
}

func (tm *Team) Pitchers() []*Player {
	return tm.pitchers
}

// Game -> Contains data on Teams and apply game functionality between them
type Game struct {
	h *Team
	a *Team
}

func (g *Game) SetHome(tm *Team) {
	g.h = tm
}

func (g *Game) SetAway(tm *Team) {
	g.a = tm
}

func (g *Game) Home() *Team {
	return g.h
}

func (g *Game) Away() *Team {
	return g.a
}
