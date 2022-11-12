package bbref

import (
	"fmt"
	"github.com/Jensen-holm/SportSimulation/bsbl"
)

func PA(h *Player, p *Player) (string, error) {
	hp := h.Probs()
	pp := p.Probs()

	outcomes := []string{"H", "BB", "HBP", "IPO"}
	weights := []float64{
		(hp["H"] + pp["H"]) / float64(2), // hit prob
		(hp["BB"] + pp["BB"]) / float64(2),
		(hp["HBP"] + pp["HBP"]) / float64(2),
		(hp["IPO"] + pp["IPO"]) / float64(2),
	}

	result, err := bsbl.Choices(outcomes, weights)
	if err != nil {
		return "", err
	}

	if result != "H" {
		return result, nil
	}
	hOutcomes := []string{"1B", "2B", "3B", "HR"}
	hWeights := []float64{
		(hp["1B"] + pp["1B"]) / float64(2),
		(hp["3B"] + pp["2B"]) / float64(2),
		(hp["3B"] + pp["3B"]) / float64(2),
		(hp["HR"] + pp["HR"]) / float64(2),
	}
	hResult, err := bsbl.Choices(hOutcomes, hWeights)
	if err != nil {
		return "", err
	}
	return hResult, nil
}

// HalfInning -> nxtHitter is the index in the lineup for the
// next hitter in the hitting team lineup
func HalfInning(nxtHitter int, hittingTm *Team, pitcher *Player) (int, int, error) {

	var outs = 0
	var ab = nxtHitter
	var runs = 0

	for outs < 3 {
		r, err := PA(hittingTm.Hitters()[ab], pitcher)
		if err != nil {
			return 0, 0, err
		}

		fmt.Println(r)
		hittingTm.Hitters()[ab].Increment(r, 1)
		pitcher.Increment(r, 1)

		if r == "IPO" || r == "SO" {
			outs += 1
		}

		ab += 1
		if ab >= len(hittingTm.Hitters()) {
			ab = 0
		}

	}
	return ab, runs, nil
}

func Inning(home, away *Team) error {

	for i := 0; i < 2; i++ {

	}

	return nil
}
