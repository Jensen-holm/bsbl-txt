package bbref

import (
	"strconv"
	"strings"
)

type Player struct {
	raw     map[string]string
	nums    map[string]int64
	attrs   map[string]string
	probs   map[string]float64
	results map[string]int
}

func NewPlayer(d map[string]string) *Player {
	np := Player{raw: d}
	np.ParseRawData(d)
	np.CalcProbs(np.Nums())
	return &np
}

// ParseRawData -> Need to make sure this function is doing what we want it to
func (p *Player) ParseRawData(d map[string]string) {
	nums := make(map[string]int64, 0)
	attrs := make(map[string]string, 0)
	for stat, val := range d {
		if n, err := strconv.ParseInt(val, 0, 64); err == nil {
			nums[stat] = n
		} else {
			attrs[stat] = val
		}
	}
	p.nums = nums
	p.attrs = attrs
}

// CalcProbs -> For some reason right now this is returning a
// map of Inf and -Inf to the probs attribute
func (p *Player) CalcProbs(n map[string]int64) {
	pr := make(map[string]float64, 0)
	isPit := strings.Contains(p.Position(), "P")

	for stat, val := range n {
		if stat == "H" || stat == "BB" || stat == "HBP" || stat == "SO" || stat == "SH" || stat == "SF" {
			if isPit {
				pr[stat] = float64(val) / float64(n["BF"])
			} else {
				pr[stat] = float64(val) / float64(n["PA"])
			}
		} else if stat == "1B" || stat == "2B" || stat == "3B" || stat == "HR" {
			if isPit {
				pr[stat] = float64(val) / float64(n["BF"])
			} else {
				pr[stat] = float64(val) / float64(n["PA"])
			}
		}
	}
	pr["IPO"] = pr["PA"] - (pr["H"] + pr["HBP"] + pr["BB"] + pr["SO"] + pr["SH"] + pr["SF"])
	p.probs = pr
}

func (p *Player) Increment(stat string, n int) {
	_, isIn := p.results[stat]
	if isIn {
		p.results[stat] += n
	} else {
		p.results[stat] = n
	}
}

func (p *Player) Raw() map[string]string {
	return p.raw
}

func (p *Player) Probs() map[string]float64 {
	return p.probs
}

func (p *Player) Nums() map[string]int64 {
	return p.nums
}

func (p *Player) Results() map[string]int {
	return p.results
}

func (p *Player) Attrs() map[string]string {
	return p.attrs
}

func (p *Player) Position() string {
	return p.attrs["Pos"]
}

func (p *Player) Name() string {
	return p.attrs["Name"]
}
