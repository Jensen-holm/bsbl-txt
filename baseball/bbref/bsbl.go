package bbref

import random "github.com/Jensen-holm/SportSimulation/random"

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

	var outs = 0
	var ab = nxtHitter
	var runs = 0

	for outs < 3 {
		r, err := PA(hittingTm.Hitters()[ab], pitcher)
		if err != nil {
			return 0, 0, err
		}

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

func Game(home, away *Team) {

	var gameOver = false
	var innings = 0.0

	for !gameOver {

		innings += .5

		if innings >= 9.5 {
			gameOver = true
		}
	}
}
