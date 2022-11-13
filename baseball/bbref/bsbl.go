package bbref

import (
	"github.com/Jensen-holm/SportSimulation/random"
)

func PA(h *Player, p *Player) (string, error) {
	hp := h.Probs()
	pp := p.Probs()

	outcomes := []string{"H", "BB", "HBP", "IPO"}
	weights := []float64{
		(hp["H"] + pp["H"]) / float64(2),
		(hp["BB"] + pp["BB"]) / float64(2),
		(hp["HBP"] + pp["HBP"]) / float64(2),
		(hp["IPO"] + pp["IPO"]) / float64(2),
	}

	result, err := random.Choices(outcomes, weights)
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
	hResult, err := random.Choices(hOutcomes, hWeights)
	if err != nil {
		return "", err
	}
	return hResult, nil
}

// HalfInning -> nxtHitter is the index in the lineup for the
// next hitter in the hitting team lineup
func HalfInning(nxtHitter int, hittingTm *Team, pitcher *Player) (int, int, error) {

	var (
		outs      = 0
		runScored = 0
		ab        = nxtHitter
		baseState = NewBaseState()
	)

	for outs < 3 {

		hitter := hittingTm.Hitters()[ab]

		r, err := PA(hitter, pitcher)
		if err != nil {
			return 0, 0, err
		}

		hitter.Increment(r, 1)
		pitcher.Increment(r, 1)

		if r == "IPO" || r == "SO" {
			outs += 1
		} else {
			runs, err := baseState.HandleBases(r)
			if err != nil {
				return 0, 0, err
			}
			hitter.Increment("RBI", runs)
			pitcher.Increment("ER", runs)
		}

		ab += 1
		if ab >= len(hittingTm.Hitters()) {
			ab = 0
		}

	}
	return ab, runScored, nil
}

func Inning(home, away *Team, hmAb, awAb int, hmPitcher, awPitcher *Player) (int, int, int, int, error) {
	nxtAbAw, ar, err := HalfInning(awAb, away, hmPitcher)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	nxtAbHm, hr, err := HalfInning(hmAb, home, awPitcher)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	return nxtAbHm, nxtAbAw, ar, hr, nil
}

// Game -> strtInn by default should be 1 when starting a new game
func Game(home, away *Team, awPitcher, hmPitcher *Player, strtInn float64) error {

	var (
		gameOver                             = false
		inning                               = strtInn
		homeScore, awayScore, homeAb, awayAb = 0, 0, 0, 0
	)

	for !gameOver {

		nxtHm, nxtAw, homeScored, awayScored, err := Inning(home, away, homeAb, awayAb, hmPitcher, awPitcher)
		if err != nil {
			return err
		}

		homeAb = nxtHm
		awayAb = nxtAw
		homeScore += homeScored
		awayScore += awayScored
		inning += .5

		if inning >= 9.5 && homeScore != awayScore {
			gameOver = true
			if homeScore > awayScore {
				home.w += 1
			} else {
				away.w += 1
			}
		} else if inning > 9.5 && homeScore == awayScore {
			// want bullpen logic in the future
			err = Game(home, away, hmPitcher, awPitcher, inning)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}
