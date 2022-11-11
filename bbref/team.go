package bbref

import (
	"sort"
)

type Team struct {
	name     string
	year     string
	hitters  []*Player
	pitchers []*Player
	lineup   []*Player
	rotation []*Player
}

func NewTeam(name string, yr string) *Team {
	return &Team{
		name: name,
		year: yr,
	}
}

// EstimateLineup -> Sorts players in the team.hitters slice by
// finding the player at each position that had the most plate appearances
func (tm *Team) EstimateLineup() {

	sort.Slice(tm.Hitters(), func(i, j int) bool {
		return tm.Hitters()[i].Nums()["PA"] > tm.Hitters()[j].Nums()["PA"]
	})

	l := make(map[string]*Player, 0)
	for _, h := range tm.Hitters() {
		pos := h.Position()
		_, isIn := l[pos]
		if !isIn {
			l[pos] = h
		} else if l[pos].Nums()["PA"] < h.Nums()["PA"] && isIn {
			l[pos] = h
		}
	}

	lineup := make([]*Player, 0)
	for _, p := range l {
		lineup = append(lineup, p)
	}
	// the slice here may be redundant but not sure
	tm.lineup = lineup[:9]
}

// EstimateRotation -> Iterates through each player that is a pitcher
// on the team and sorts them by batters faced, returns the top 5
// as an estimation of who on the team pitched the most
func (tm *Team) EstimateRotation() {
	r := make([]*Player, 0)
	// BF may not be the best way to do this
	sort.Slice(r, func(i, j int) bool {
		return r[i].Attrs()["BF"] > r[j].Attrs()["BF"]
	})
	tm.rotation = r[:5]
}

// may not need some of these

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

func (tm *Team) Lineup() []*Player {
	return tm.lineup
}

func (tm *Team) Rotation() []*Player {
	return tm.rotation
}
