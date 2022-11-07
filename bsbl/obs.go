package bsbl

import "strconv"

type Hitter struct {
	name string
	pos  string
	age  string
	PA   int64
	AB   int64
	H    int64
	B2   int64
	B3   int64
	HR   int64
	SB   int64
	CS   int64
	BB   int64
	SO   int64
	TB   int64
	GDP  int64
	HBP  int64
	SH   int64
	SF   int64
	IBB  int64
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
	tb, _ := strconv.ParseInt(d["TB"], 0, 64)
	gdp, _ := strconv.ParseInt(d["GDP"], 0, 64)
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
		H:    h,
		B2:   b2,
		B3:   b3,
		HR:   hr,
		SB:   sb,
		CS:   cs,
		BB:   bb,
		SO:   so,
		TB:   tb,
		GDP:  gdp,
		HBP:  hbp,
		SH:   sh,
		SF:   sf,
		IBB:  ibb,
	}

}

type Pitcher struct {
	name string
	pos  string
	age  string
	HA   int64
	B2   int64
	B3   int64
	HRA  int64
	BB   int64
	SO   int64
	WP   int64
	BK   int64
	BF   int64
	ER   int64
	IP   float64
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
	wp, _ := strconv.ParseInt(d["WP"], 0, 64)
	bk, _ := strconv.ParseInt(d["BK"], 0, 64)
	bf, _ := strconv.ParseInt(d["BF"], 0, 64)
	er, _ := strconv.ParseInt(d["ER"], 0, 64)
	ip, _ := strconv.ParseFloat(d["IP"], 64)
	return &Pitcher{
		name: d["Name"],
		pos:  d["Pos"],
		age:  d["Age"],
		HA:   ha,
		B2:   b2,
		B3:   b3,
		HRA:  hr,
		BB:   bb,
		SO:   so,
		WP:   wp,
		BK:   bk,
		BF:   bf,
		ER:   er,
		IP:   ip,
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
