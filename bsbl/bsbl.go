package bsbl

//// Choices -> Select an element from a slice given weights
//// assumes that the length of the sl and weights slices is equal
//// and that the weights are floats between 0 and 1
//func Choices(sl []interface{}, weights []float64) string {
//	var max int = 0
//	var c string
//	for i := 0; i < len(sl); i++ {
//		rNum := rand.Intn(100)
//		if rNum <= int((100 * weights[i])) {
//			return sl[i]
//		} else {
//			if int(weights[i]) > max {
//				max = sl[i]
//			}
//		}
//	}
//	return sl[max]
//}
//
//func PA(h *Player, p *Player) (string, error) {
//	// choose an outcome weighted by probabilities
//	hProbs := h.
//
//	return "", nil
//}
