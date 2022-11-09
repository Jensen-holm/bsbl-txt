package bbref

import (
	"sort"
	"strconv"
)

// sadly we may have to create more objects to accommodate data
// scraped from other sites like baseball savant, probably want to
// split it all into their own packages when we get to that point

// Hitter -> this is specific to bbref players
type Hitter struct {
	raw   map[string]string
	nums  map[string]int64
	attrs map[string]string
	probs map[string]float64
}

func (h *Hitter) Name() string { return h.Attrs()["Name"] }

func (h *Hitter) Stats() map[string]int64 { return h.nums }

func (h *Hitter) Attrs() map[string]string { return h.attrs }

func (h *Hitter) Probs() map[string]float64 { return h.probs }

func (h *Hitter) ParseStats(d map[string]string) {
	statMap := make(map[string]int64, 0)
	attrMap := make(map[string]string, 0)
	for stat, val := range d {
		if s, err := strconv.ParseInt(val, 0, 64); err == nil {
			statMap[stat] = s
		} else {
			attrMap[stat] = val
		}
	}
	h.nums = statMap
	h.attrs = attrMap
}

func (h *Hitter) CalcProbs(n map[string]int64) {
	p := make(map[string]float64, 0)
	for stat, val := range n {
		if stat == "H" || stat == "BB" || stat == "HBP" || stat == "SO" || stat == "SH" || stat == "SF" {
			p[stat] = float64(val) / float64(n["PA"])
			continue
		}
		if stat == "1B" || stat == "2B" || stat == "3B" || stat == "HR" {
			p[stat] = float64(val) / float64(n["PA"])
			continue
		}
		if stat == "SB" {
			p[stat] = float64(val)
			p["ATT"] = float64(val) + float64(n["CS"])
			continue
		}
	}
	p["IPO"] = p["PA"] - (p["H"] + p["HBP"] + p["BB"] + p["SO"] + p["SH"] + p["SF"])
	h.probs = p
}

// NewHitter -> Constructor for hitter objects
func NewHitter(d map[string]string) *Hitter {
	h := new(Hitter)
	h.ParseStats(d)
	h.CalcProbs(h.nums)
	return h
}

type Pitcher struct {
	raw   map[string]string
	nums  map[string]int64
	attrs map[string]string
	probs map[string]float64
}

// Duplicated code below but go does not support inheritance

func (p *Pitcher) Name() string { return p.Attrs()["Name"] }

func (p *Pitcher) Stats() map[string]int64 { return p.nums }

func (p *Pitcher) Attrs() map[string]string { return p.attrs }

func (p *Pitcher) Probs() map[string]float64 { return p.probs }

func (p *Pitcher) ParseStats(d map[string]string) {
	statMap := make(map[string]int64, 0)
	attrMap := make(map[string]string, 0)
	for stat, val := range d {
		if s, err := strconv.ParseInt(val, 0, 64); err == nil {
			statMap[stat] = s
		} else {
			attrMap[stat] = d[stat]
		}
	}
	p.nums = statMap
	p.attrs = attrMap
}

func (p *Pitcher) CalcProbs(n map[string]int64) {
	pr := make(map[string]float64, 0)
	for stat, val := range n {
		if stat == "H" || stat == "BB" || stat == "HBP" || stat == "SO" || stat == "SH" || stat == "SF" {
			pr[stat] = float64(val) / float64(n["PA"])
			continue
		}
		if stat == "1B" || stat == "2B" || stat == "3B" || stat == "HR" {
			pr[stat] = float64(val) / float64(n["PA"])
			continue
		}
	}
	pr["IPO"] = pr["PA"] - (pr["H"] + pr["HBP"] + pr["BB"] + pr["SO"] + pr["SH"] + pr["SF"])
	p.probs = pr
}

func NewPitcher(d map[string]string) *Pitcher {
	p := new(Pitcher)
	p.ParseStats(d)
	p.CalcProbs(p.Stats())
	return p
}

// Team -> Contains Data about a team
type Team struct {
	name     string
	year     string
	hitters  []*Hitter
	pitchers []*Pitcher
	lineup   []*Hitter
	rotation []*Pitcher
}

func (tm *Team) SetName(n string) { tm.name = n }

func (tm *Team) SetYear(yr string) { tm.year = yr }

func (tm *Team) SetHitters(hitters []*Hitter) { tm.hitters = hitters }

func (tm *Team) SetPitchers(pitchers []*Pitcher) { tm.pitchers = pitchers }

func (tm *Team) Year() string { return tm.year }

func (tm *Team) Name() string { return tm.name }

func (tm *Team) Hitters() []*Hitter { return tm.hitters }

func (tm *Team) Pitchers() []*Pitcher { return tm.pitchers }

// EstimateLineup -> Sorts players in the team.hitters slice by
// finding the player at each position that had the most plate appearances
func (tm *Team) EstimateLineup() {
	l := make(map[string]*Hitter, 0)
	for _, h := range tm.Hitters() {
		// if the position already exists in the map
		if _, ok := l[h.Attrs()["Pos"]]; ok && h.Stats()["PA"] > 0 {
			// check if the h has a higher prob than the one already in there
			l[h.Attrs()["Pos"]] = h
			continue
		}
		// if it doesn't exist already, make them the defualt
		l[h.Attrs()["Pos"]] = h
	}
	line := make([]*Hitter, 0)
	for _, h := range l {
		line = append(line, h)
	}
	sort.Slice(line, func(i, j int) bool {
		return line[i].Probs()["H"] > line[j].Probs()["H"]
	})
	tm.lineup = line[:9]
}

func (tm *Team) EstimateRotation() {
	r := make([]*Pitcher, 0)
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

func (tm *Team) Lineup() []*Hitter { return tm.lineup }

func (tm *Team) Rotation() []*Pitcher { return tm.rotation }
