package random

import (
	"fmt"
	"math/rand"
	"sort"
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

	// sort the weights from smallest to largest
	sort.Slice(weights, func(i, j int) bool {
		return weights[i] < weights[j]
	})

	rNum := rand.Intn(100)

	for i := 0; i < l; i++ {
		wt := weights[i]
		if float64(rNum) <= (wt * 100) {
			return arr[i], nil
		}
	}
	// returns most probable if the random num wasn't smaller than any of the probabilities
	return arr[len(arr)-1], nil
}
