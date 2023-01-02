package bbref

import (
	"sort"
	"strings"
)

type Team struct {
	Name     string
	Year     string
	Hitters  []*Player
	Pitchers []*Player
	Lineup   []*Player
	Rotation []*Player
	Stats    map[string]any
	W        int
}

func NewTeam(input string) *Team {
	spl := strings.Split(input, " ")
	return &Team{
		Name: strings.Title(strings.Join(spl[1:], " ")),
		Year: spl[0],
	}
}

// EstimateLineup -> Sorts players in the team.hitters slice by
// finding the player at each position that had the most plate appearances
func (tm *Team) EstimateLineup() {

	hits := tm.Hitters

	sort.Slice(hits, func(i, j int) bool {
		return hits[i].Nums()["PA"] > hits[j].Nums()["PA"]
	})

	l := make(map[string]*Player, 0)
	for _, h := range tm.Hitters {
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
	tm.Lineup = lineup[:9]
}

// EstimateRotation -> Iterates through each player that is a pitcher
// on the team and sorts them by batters faced, returns the top 5
// as an estimation of who on the team pitched the most
func (tm *Team) EstimateRotation() {
	sort.Slice(tm.Pitchers, func(i, j int) bool {
		return tm.Pitchers[i].Attrs()["BF"] > tm.Pitchers[j].Attrs()["BF"]
	})
	tm.Rotation = tm.Pitchers[:5]
}

func (tm *Team) SetName(n string) {
	tm.Name = n
}

func (tm *Team) SetYear(yr string) {
	tm.Year = yr
}

func (tm *Team) SetHitters(hitters []*Player) {
	tm.Hitters = hitters
}

func (tm *Team) SetPitchers(pitchers []*Player) {
	tm.Pitchers = pitchers
}
