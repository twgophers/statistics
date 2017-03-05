package statistics

import (
	"math"
)

func validate(sample []float64) {
	if len(sample) == 0 {
		panic("empty sample supplyed")
	}
}

type binaryCondition func(v1, v2 float64) float64

func matchingValue(fn binaryCondition, initial float64, sample []float64) float64 {
	validate(sample)

	current := initial
	for _, value := range sample {
		current = fn(current, value)
	}
	return current
}

func MaxIn(sample []float64) float64 {
	return matchingValue(math.Max, math.Inf(-1), sample)
}

func MinIn(sample []float64) float64 {
	return matchingValue(math.Min, math.Inf(+1), sample)
}
