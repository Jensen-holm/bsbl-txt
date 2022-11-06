package bsbl

type Player struct {
	data map[string]string
}

type Team struct {
	name     string
	year     string
	hitters  []Player
	pitchers []Player
}

func (tm *Team) SetHitters(hitters []Player) {
	tm.hitters = hitters
}

func (tm *Team) SetPitchers(pitchers []Player) {
	tm.pitchers = pitchers
}
