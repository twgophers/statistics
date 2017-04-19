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

func (sample Sample) Quantile(percentile float64) float64 {
	pIndex := int(percentile * float64(sample.size()))

	sort.Float64s(sample)

	return sample[pIndex]
}

func (sample Sample) Mode() []float64 {
	sample.check()

	counts := count(sample)

	maxQuantitie := maxValue(counts)

	modes := []float64{}

	for k, v := range counts {
		if v == maxQuantitie {
			modes = append(modes, k)
		}
	}

	return modes
}

func maxValue(counts map[float64]int64) int64 {
	var quantities Sample = make([]float64, 0, len(counts))
	for _, v := range counts {
		quantities = append(quantities, float64(v))
	}

	return int64(quantities.Max())
}

func count(sample Sample) map[float64]int64 {
	counts := map[float64]int64{}

	for _, value := range sample {
		v, ok := counts[value]
		if !ok {
			v = 0
		}
		counts[value] = v + 1
	}

	return counts
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
