package bsbl

type Player struct {
	data map[string]string
}

func (p *Player) SetData(d map[string]string) {
	p.data = d
}

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

func (tm *Team) GetYear() string {
	return tm.year
}

func (tm *Team) GetName() string {
	return tm.name
}
