package statistics

import (
	"fmt"
	"math"
	"sort"

	"github.com/twgophers/collections"
	"github.com/twgophers/linalg"
)

func Sum(sample collections.Vector) float64 {
	total := 0.0
	for _, value := range sample {
		total += value
	}
	return total
}

func Mean(sample collections.Vector) float64 {
	check(sample)

	return Sum(sample) / float64(sample.Len())
}

func Median(sample collections.Vector) float64 {
	check(sample)

	sort.Float64s(sample)

	half := sample.Len() / 2

	if oddSize(sample) {
		return sample[half]
	}

	return Mean(collections.Vector{sample[half-1], sample[half]})
}

func Quantile(sample collections.Vector, percentile float64) float64 {
	pIndex := int(percentile * float64(sample.Len()))

	sort.Float64s(sample)

	return sample[pIndex]
}

func Mode(sample collections.Vector) collections.Vector {
	check(sample)
	counter := collections.NewCounter(sample)
	maxQuantity := counter.MaxValue()

	modes := make(collections.Vector, 0, len(sample))

	for k, v := range counter.Items {
		if v == maxQuantity {
			modes = append(modes, k)
		}
	}

	sort.Sort(sort.Reverse(modes))

	return modes
}

func DataRange(sample collections.Vector) float64 {
	return sample.Max() - sample.Min()
}

func DispersionMean(sample collections.Vector) collections.Vector {
	mean := Mean(sample)
	dispersion := make(collections.Vector, 0, cap(sample))

	for _, value := range sample {
		dispersion = append(dispersion, value-mean)
	}

	return dispersion
}

func Variance(sample collections.Vector) float64 {
	checkMinimumSize(sample.Len(), 1)
	dispersionMean := DispersionMean(sample)

	return linalg.SumOfSquares(dispersionMean) / float64(sample.Len()-1)
}

func StandardDeviation(sample collections.Vector) float64 {
	return math.Sqrt(Variance(sample))
}

func InterQuantileRange(sample collections.Vector) float64 {
	return Quantile(sample, 0.75) - Quantile(sample, 0.25)
}

func Covariance(x, y collections.Vector) float64 {
	n := x.Len()
	checkMinimumSize(n, 1)
	return linalg.Dot(DispersionMean(x), DispersionMean(y)) / float64(n-1)
}

func Correlation(x, y collections.Vector) float64 {
	standardDeviationX := StandardDeviation(x)
	standardDeviationY := StandardDeviation(y)
	if standardDeviationX > 0 && standardDeviationY > 0 {
		return Covariance(x, y) / standardDeviationX / standardDeviationY
	}
	return float64(0)
}

func checkMinimumSize(value, minimum int) {
	if value <= minimum {
		panic(fmt.Errorf("The minimum size was not obeyed - %d", minimum))
	}
}

func oddSize(sample collections.Vector) bool {
	return sample.Len()%2 == 1
}

func check(sample collections.Vector) {
	if sample.Empty() {
		panic("Operation Not allowed with empty sample")
	}
}
