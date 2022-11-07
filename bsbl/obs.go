package bsbl

import (
	"strconv"
)

type Hitter struct {
	name string
	pos  string
	age  string
	PA   int64
	AB   int64
	H    float64
	B1   float64
	B2   float64
	B3   float64
	HR   float64
	SB   float64
	// ATT is the probability of attempting to steal a base when on first
	ATT float64
	BB  float64
	SO  float64
	HBP float64
	SH  float64
	SF  float64
	IBB float64
	IPO float64
}

// NewHitter -> Assumes that the
func NewHitter(d map[string]string) *Hitter {
	pa, _ := strconv.ParseInt(d["PA"], 0, 64)
	ab, _ := strconv.ParseInt(d["AB"], 0, 64)
	h, _ := strconv.ParseInt(d["H"], 0, 64)
	b2, _ := strconv.ParseInt(d["2B"], 0, 64)
	b3, _ := strconv.ParseInt(d["3B"], 0, 64)
	hr, _ := strconv.ParseInt(d["HR"], 0, 64)
	so, _ := strconv.ParseInt(d["SO"], 0, 64)
	hbp, _ := strconv.ParseInt(d["HBP"], 0, 64)
	sh, _ := strconv.ParseInt(d["SH"], 0, 64)
	sf, _ := strconv.ParseInt(d["SF"], 0, 64)
	ibb, _ := strconv.ParseInt(d["IBB"], 0, 64)
	sb, _ := strconv.ParseInt(d["SB"], 0, 64)
	cs, _ := strconv.ParseInt(d["CS"], 0, 64)
	bb, _ := strconv.ParseInt(d["BB"], 0, 64)
	return &Hitter{
		name: d["Name"],
		pos:  d["Pos"],
		age:  d["Age"],
		PA:   pa,
		AB:   ab,
		H:    float64(h) / float64(pa),
		B1:   float64(h-(b2+b3+hr)) / float64(pa),
		B2:   float64(b2) / float64(h),
		B3:   float64(b3) / float64(h),
		HR:   float64(hr) / float64(h),
		SB:   float64(sb) / float64(cs),
		ATT:  float64(sb) / float64(bb+(h-(b2+b3+hr))),
		BB:   float64(bb) / float64(pa),
		SO:   float64(so) / float64(pa),
		HBP:  float64(hbp) / float64(pa),
		SH:   float64(sh) / float64(pa),
		SF:   float64(sf) / float64(pa),
		IBB:  float64(ibb) / float64(pa),
		// double check this one
		IPO: float64(ab-h-bb-hbp) / float64(pa),
	}
}

type Pitcher struct {
	name string
	pos  string
	age  string
	HA   float64
	B1   float64
	B2   float64
	B3   float64
	HRA  float64
	BB   float64
	SO   float64
	BK   float64
	BF   int64
	IP   float64
	HBP  float64
	IPO  float64
}

// NewPitcher -> Assumes that the map entered into this function is that
// of a player scraped off of a baseball reference team_pitching table
// I wish I knew a better way to write these functions
func NewPitcher(d map[string]string) *Pitcher {
	ha, _ := strconv.ParseInt(d["H"], 0, 64)
	b2, _ := strconv.ParseInt(d["2B"], 0, 64)
	b3, _ := strconv.ParseInt(d["3B"], 0, 64)
	hr, _ := strconv.ParseInt(d["HR"], 0, 64)
	bb, _ := strconv.ParseInt(d["BB"], 0, 64)
	so, _ := strconv.ParseInt(d["SO"], 0, 64)
	bf, _ := strconv.ParseInt(d["BF"], 0, 64)
	ip, _ := strconv.ParseFloat(d["IP"], 64)
	hbp, _ := strconv.ParseInt(d["HBP"], 0, 64)
	return &Pitcher{
		name: d["Name"],
		pos:  d["Pos"],
		age:  d["Age"],
		HA:   float64(ha) / float64(bf),
		B1:   float64(ha-(b2+b3+hr)) / float64(ha),
		B2:   float64(b2) / float64(ha),
		B3:   float64(b3) / float64(ha),
		HRA:  float64(hr) / float64(ha),
		BB:   float64(bb) / float64(bf),
		SO:   float64(so) / float64(bf),
		BF:   bf,
		IP:   ip,
		HBP:  float64(hbp) / float64(bf),
		IPO:  float64(bf-bb-so-ha-hbp) / float64(bf),
	}
}

// Team -> Contains Data about a team
type Team struct {
	name     string
	year     string
	hitters  []*Hitter
	pitchers []*Pitcher
}

func (tm *Team) SetName(n string) { tm.name = n }

func (tm *Team) SetYear(yr string) { tm.year = yr }

func (tm *Team) SetHitters(hitters []*Hitter) { tm.hitters = hitters }

func (tm *Team) SetPitchers(pitchers []*Pitcher) { tm.pitchers = pitchers }

func (tm *Team) Year() string { return tm.year }

func (tm *Team) Name() string { return tm.name }

func (tm *Team) Hitters() []*Hitter { return tm.hitters }

func (tm *Team) Pitchers() []*Pitcher { return tm.pitchers }
