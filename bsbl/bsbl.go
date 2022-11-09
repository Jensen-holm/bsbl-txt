package bsbl

func PA(h *Hitter, p *Pitcher) (string, error) {
	hp := h.Probs()
	pp := p.Probs()

	outcomes := []string{"H", "BB", "HBP", "IPO"}
	weights := []float64{
		(hp["H"] + pp["H"]) / float64(2), // hit prob
		(hp["BB"] + pp["BB"]) / float64(2),
		(hp["HBP"] + pp["HBP"]) / float64(2),
		(hp["IPO"] + pp["IPO"]) / float64(2),
	}

	result, err := Choices(outcomes, weights)
	if err != nil {
		return "", err
	}

	if result == "H" {
		hOutcomes := []string{"1B", "2B", "3B", "HR"}
		hWeights := []float64{
			(hp["1B"] + pp["1B"]) / float64(2),
			(hp["3B"] + pp["2B"]) / float64(2),
			(hp["3B"] + pp["3B"]) / float64(2),
			(hp["HR"] + pp["HR"]) / float64(2),
		}
		hResult, err := Choices(hOutcomes, hWeights)
		if err != nil {
			return "", err
		}
		return hResult, nil
	}
	return result, nil
}

// HalfInning -> nxtHitter is the index in the lineup for the
// next hitter in the hitting team lineup
func HalfInning(nxtHitter int, hittingTm *Team, pitcher *Pitcher) (int, int, error) {

	var outs = 0
	var ab = nxtHitter
	var runs = 0

	for {

		r, err := PA(hittingTm.Hitters()[ab], pitcher)
		if err != nil {
			return 0, 0, err
		}

		if r == "IPO" || r == "SO" {
			outs += 1
		}

		ab += 1
		if ab >= len(hittingTm.Hitters()) {
			ab = 0
		}

		if outs > 2 {
			break
		}

	}
	return ab, runs, nil
}
