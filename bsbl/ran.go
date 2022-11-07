package bsbl

import (
	"fmt"
	"math/rand"
)

// Creating some functions not specific to baseball
// so that we can drive the simulation based on weighted
// probabilities monte carlo simulation no matter which
// metrics we decide to build the simulation around

// Choices -> Chooses an element from a slice of strings based on weighted
// requires that the length of the two input parameters are equal
func Choices(arr []string, weights []float64) (string, error) {
	l := len(arr)
	if l != len(weights) {
		return "", fmt.Errorf("in the choices function the length of the outcome and weight slices must be equal")
	}

	var max int
	for i := 0; i < l; i++ {

		rNum := rand.Intn(100)
		wt := int(100 * weights[i])

		if rNum <= wt {
			return arr[i], nil
		}

		if int(weights[i]) >= max {
			max = i
		}
	}

	return arr[max], nil
}
