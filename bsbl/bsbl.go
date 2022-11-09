package bsbl

func PA(h *Hitter, p *Pitcher) (string, error) {

	outcomes := []string{"H", "BB", "HBP", "IPO"}
	weights := []float64{
		(h.H + p.HA) / float64(2), // hit prob
		(h.BB + p.BB) / float64(2),
		(h.HBP + p.HBP) / float64(2),
		(h.IPO + p.IPO) / float64(2),
	}

	result, err := Choices(outcomes, weights)
	if err != nil {
		return "", err
	}

	if result == "H" {
		hOutcomes := []string{"1B", "2B", "3B", "HR"}
		hWeights := []float64{
			(h.B1 + p.B1) / float64(2),
			(h.B2 + p.B2) / float64(2),
			(h.B3 + p.B3) / float64(2),
			(h.HR + p.HRA) / float64(2),
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
