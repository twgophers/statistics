package statistics

import (
	"math"
	"sort"
)

func validate(sample Sample) {
	if len(sample) == 0 {
		panic("empty sample supplyed")
	}
}

type binaryCondition func(v1, v2 float64) float64
type Sample []float64

func matchingValue(fn binaryCondition, initial float64, sample Sample) float64 {
	validate(sample)

	current := initial
	for _, value := range sample {
		current = fn(current, value)
	}
	return current
}

func (sample Sample) Max() float64 {
	return matchingValue(math.Max, math.Inf(-1), sample)
}

func (sample Sample) Min() float64 {
	return matchingValue(math.Min, math.Inf(+1), sample)
}

func (sample Sample) Sum() float64 {
	total := 0.0
	for _, value := range sample {
		total += value
	}
	return total
}

func (sample Sample) Mean() float64 {
	sample.check()

	return sample.Sum() / float64(sample.size())
}

func (sample Sample) Median() float64 {
	sample.check()

	sort.Float64s(sample)

	half := sample.size() / 2

	if sample.oddSize() {
		return sample[half]
	}

	return Sample{sample[half-1], sample[half]}.Mean()
}

func (sample Sample) size() int {
	return len(sample)
}

func (sample Sample) empty() bool {
	return sample.size() == 0
}

func (sample Sample) oddSize() bool {
	return sample.size()%2 == 1
}

func (sample Sample) check() {
	if sample.empty() {
		panic("Operation Not allowed with empty sample")
	}
}
