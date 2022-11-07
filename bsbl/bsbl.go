package bsbl

func PA(h *Hitter, p *Pitcher) {

	outcomes := []string{"H", "BB", "HBP", "IPO"}
	weights := []float64{
		(h.H + p.HA) / float64(2), // hit prob
		(h.BB + p.BB) / float64(2),
		(h.HBP + p.HBP) / float64(2),
		(h.IPO + p.IPO) / float64(2),
	}

	result, err := Choices(outcomes, weights)
	if err != nil {
		return
	}

	if result == "H" {
		hOutcomes := []string{"1B", "2B", "3B", "HR"}
		hWeights := []float64{
			(h.B1 + p.B1) / float64(2),
			(h.B2 + p.B2) / float64(2),
			(h.B3 + p.B3) / float64(2),
			(h.HR + p.HRA) / float64(2),
		}
		_, err := Choices(hOutcomes, hWeights)
		if err != nil {
			return
		}
		return
	}
	return
}
