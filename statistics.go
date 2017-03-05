package statistics

import (
	"math"
)

func MaxIn(sample []float64) float64 {
	if len(sample) == 0 {
		panic("empty sample supplyed")
	}

	currentMax := math.Inf(-1)
	for _, value := range sample {
		currentMax = math.Max(currentMax, value)
	}
	return currentMax
}
