package bbref

import (
	"sort"
)

// Team -> Contains Data about a team
type Team struct {
	name     string
	year     string
	hitters  []*Player
	pitchers []*Player
	lineup   []*Player
	rotation []*Player
}

func (tm *Team) SetName(n string) { tm.name = n }

func (tm *Team) SetYear(yr string) { tm.year = yr }

func (tm *Team) SetHitters(hitters []*Player) { tm.hitters = hitters }

func (tm *Team) SetPitchers(pitchers []*Player) { tm.pitchers = pitchers }

func (tm *Team) Year() string { return tm.year }

func (tm *Team) Name() string { return tm.name }

func (tm *Team) Hitters() []*Player { return tm.hitters }

func (tm *Team) Pitchers() []*Player { return tm.pitchers }

// EstimateLineup -> Sorts players in the team.hitters slice by
// finding the player at each position that had the most plate appearances
func (tm *Team) EstimateLineup() {
	l := make(map[string]*Player, 0)
	for _, h := range tm.Hitters() {
		// if the position already exists in the map
		if _, ok := l[h.Attrs()["Pos"]]; ok && h.Nums()["PA"] > 0 {
			// check if the h has a higher prob than the one already in there
			l[h.Attrs()["Pos"]] = h
			continue
		}
		// if it doesn't exist already, make them the defualt
		l[h.Attrs()["Pos"]] = h
	}
	line := make([]*Player, 0)
	for _, h := range l {
		line = append(line, h)
	}
	sort.Slice(line, func(i, j int) bool {
		return line[i].Probs()["H"] > line[j].Probs()["H"]
	})
	tm.lineup = line[:9]
}

func (tm *Team) EstimateRotation() {
	r := make([]*Player, 0)
	for _, p := range tm.Pitchers() {
		if p.Attrs()["Pos"] != "SP" {
			continue
		}
		r = append(r, p)
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].Attrs()["BF"] > r[j].Attrs()["BF"]
	})
	tm.rotation = r[:5]
}

func (tm *Team) Lineup() []*Player { return tm.lineup }

func (tm *Team) Rotation() []*Player { return tm.rotation }
