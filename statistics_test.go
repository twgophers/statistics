package statistics

import (
	"reflect"
	"testing"

	"github.com/twgophers/collections"
)

func TestSum(t *testing.T) {
	cases := []struct {
		sample collections.Vector
		want   float64
	}{
		{collections.Vector{7.0}, 7.0},
		{collections.Vector{32.0, 7.0}, 39.0},
		{collections.Vector{}, 0.0},
	}
	for _, c := range cases {
		gotSum := Sum(c.sample)
		if gotSum != c.want {
			t.Errorf("Expected total (%v) summing up (%v) but got (%v)", c.want, c.sample, gotSum)
		}
	}
}

func TestMean(t *testing.T) {
	cases := []struct {
		sample collections.Vector
		want   float64
	}{
		{collections.Vector{7.0}, 7.0},
		{collections.Vector{13.0, 14.0}, 13.5},
	}
	for _, c := range cases {
		gotMean := Mean(c.sample)

		if gotMean != c.want {
			t.Errorf("Expected mean of (%v) for (%v) but got (%v)", c.want, c.sample, gotMean)
		}
	}
}

func TestMeanPanicsWhenEmptySlice(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected mean panic when empty sample")
		}
	}()

	Mean(collections.Vector{})
}

func TestMedian(t *testing.T) {
	cases := []struct {
		sample collections.Vector
		want   float64
	}{
		{collections.Vector{7.0}, 7.0},
		{collections.Vector{8.0, 11.0}, 9.5},
		{collections.Vector{7.0, 8.0, 11.0}, 8.0},
		{collections.Vector{7.0, 9.0, 10.0, 17.0}, 9.5},
		{collections.Vector{7.0, 10.0, 17.0, 9.0}, 9.5},
	}
	for _, c := range cases {
		gotMedian := Median(c.sample)

		if gotMedian != c.want {
			t.Errorf("Expected median (%v) for (%v) but got (%v)", c.want, c.sample, gotMedian)
		}
	}
}

func TestMedianPanicsWhenEmptySlice(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected median panic when empty sample")
		}
	}()

	Median(collections.Vector{})
}

func TestQuantile(t *testing.T) {
	cases := []struct {
		sample     collections.Vector
		percentile float64
		want       float64
	}{
		{collections.Vector{7.0}, 0.99, 7.0},
		{collections.Vector{7.0, 9.0, 10.0, 13.0, 17.0}, 0.75, 13.0},
		{collections.Vector{7.0, 9.0, 13.0, 10.0, 17.0}, 0.75, 13.0},
	}

	for _, c := range cases {
		gotQuantile := Quantile(c.sample, c.percentile)
		if gotQuantile != c.want {
			t.Errorf("The expected quantile for (%v) with percentile of (%.2f) was (%.2f) but got (%.2f)", c.sample, c.percentile, c.want, gotQuantile)
		}
	}
}

func TestMode(t *testing.T) {
	cases := []struct {
		sample collections.Vector
		want   collections.Vector
	}{
		{collections.Vector{7.0}, collections.Vector{7.0}},
		{collections.Vector{7.0, 13.0, 13.0}, collections.Vector{13.0}},
		{collections.Vector{17.0, 7.0, 13.0, 17.0, 13.0}, collections.Vector{17.0, 13.0}},
	}

	for _, c := range cases {
		gotMode := Mode(c.sample)

		if !reflect.DeepEqual(gotMode, c.want) {
			t.Errorf("Expected mode (%v) for (%v) but got (%v)", c.want, c.sample, gotMode)
		}
	}
}

func TestModeFailWhenEmptySample(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("A panic was expected but nothing happened when calculate mode for empty Sample")
		}
	}()

	Mode(collections.Vector{})
}

func TestDataRange(t *testing.T) {
	cases := []struct {
		sample collections.Vector
		want   float64
	}{
		{collections.Vector{10.0, 11.0, 12.0}, 2.0},
		{collections.Vector{0.0, 11.0, 12.0}, 12.0},
		{collections.Vector{-1.0, 11.0, 12.0}, 13.0},
		{collections.Vector{1.0, 1.0, 1.0}, 0.0},
		{collections.Vector{11.0}, 0.0},
	}

	for _, c := range cases {
		gotDataRange := DataRange(c.sample)
		if !reflect.DeepEqual(gotDataRange, c.want) {
			t.Errorf("Expected DataRange (%v) for (%v) but got (%v)", c.want, c.sample,
				gotDataRange)
		}
	}
}

func TestDataRangeWhenSampleIsEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("A panic was expected but nothing happened when calculate DataRange for empty Sample")
		}
	}()

	DataRange(collections.Vector{})
}

func TestDispersionMean(t *testing.T) {
	cases := []struct {
		sample collections.Vector
		want   collections.Vector
	}{
		{collections.Vector{10.0, 20.0, 30.0}, collections.Vector{-10.0, 0.0, 10.0}},
		{collections.Vector{0.0, 10.0, 20.0}, collections.Vector{-10.0, 0.0, 10.0}},
		{collections.Vector{0.0, 0.0, 0.0}, collections.Vector{0.0, 0.0, 0.0}},
		{collections.Vector{-1.0, 0.0, 1.0}, collections.Vector{-1.0, 0.0, 1.0}},
		{collections.Vector{0.0, 20.0, 30.0}, collections.Vector{-16.666666666666668, 3.333333333333332, 13.333333333333332}},
	}

	for _, c := range cases {
		gotDispersionMean := DispersionMean(c.sample)
		if !reflect.DeepEqual(gotDispersionMean, c.want) {
			t.Errorf("Expected DispersionMean (%v) for (%v) but got (%v)", c.want, c.sample,
				gotDispersionMean)
		}
	}
}

func TestDispersionMeanWhenSampleIsEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("A panic was expected but nothing happened when calculate DispersionMean for empty Sample")
		}
	}()

	DispersionMean(collections.Vector{})
}

func TestVariance(t *testing.T) {
	cases := []struct {
		sample collections.Vector
		want   float64
	}{
		{collections.Vector{1.0, 2.0}, 0.5},
		{collections.Vector{1.0, 2.0, 3.0}, 1.0},
	}

	for _, c := range cases {
		gotVariance := Variance(c.sample)
		if gotVariance != c.want {
			t.Errorf("Expected Variance (%v) for (%v) but got (%v)", c.want, c.sample,
				gotVariance)
		}
	}
}

func TestVariance_WhenVectorIsEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("A panic was expected but nothing happened when calculate Variance for empty Sample")
		}
	}()

	Variance(collections.Vector{})
}

func TestVariance_WhenVectorHasOneElement(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("A panic was expected but nothing happened when calculate Variance for empty Sample")
		}
	}()

	Variance(collections.Vector{1.0})
}

func TestStandardDeviation(t *testing.T) {
	cases := []struct {
		sample collections.Vector
		want   float64
	}{
		{collections.Vector{1.0, 2.0, 3.0}, 1.0},
		{collections.Vector{1.0, 2.0}, 0.7071067811865476},
	}

	for _, c := range cases {
		got := StandardDeviation(c.sample)
		if got != c.want {
			t.Errorf("StandarDeviation(%v) want: %v but got: %v",
				c.sample, c.want, got)
		}
	}
}

func TestStandardDeviation_WhenVectorHasOneElement(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("A panic was expected but nothing happened when calculate Variance for empty Sample")
		}
	}()

	StandardDeviation(collections.Vector{1.0})
}

func TestInterQuantileRange(t *testing.T) {
	cases := []struct {
		sample collections.Vector
		want   float64
	}{
		{collections.Vector{1.0, 2.0, 3.0}, 2.0},
		{collections.Vector{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}, 3.0},
	}
	for _, c := range cases {
		got := InterQuantileRange(c.sample)
		if got != c.want {
			t.Errorf("InterQuantileRange(%v) want: %v but got: %v",
				c.sample, c.want, got)
		}
	}
}

func TestInterQuantileRange_WhenVectorHasOneElement(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("A panic was expected but nothing happened when calculate Variance for empty Sample")
		}
	}()

	InterQuantileRange(collections.Vector{})
}

func TestCovariance(t *testing.T) {
	cases := []struct {
		x, y collections.Vector
		want float64
	}{
		{collections.Vector{1.0, 2.0}, collections.Vector{1.0, 2.0}, 0.5},
		{collections.Vector{1.0, 2.0, 3.0}, collections.Vector{1.0, 2.0}, 0.25},
		{collections.Vector{1.0, 1.0}, collections.Vector{1.0, 1.0}, 0.0},
		{collections.Vector{1.0, 1.0}, collections.Vector{1.0}, 0.0},
	}

	for _, c := range cases {
		gotVariance := Covariance(c.x, c.y)
		if gotVariance != c.want {
			t.Errorf("Expected Covariance (%v, %v) want: (%v) but got: (%v)",
				c.x, c.y, c.want, gotVariance)
		}
	}
}

func TestCovariance_WhenVectorHasOneElement(t *testing.T) {
	cases := []struct {
		x, y collections.Vector
	}{
		{collections.Vector{1.0}, collections.Vector{1.0, 2.0, 3.0}},
		{collections.Vector{}, collections.Vector{1.0, 2.0, 3.0}},
	}

	for _, c := range cases {
		defer func() {
			if recover() == nil {
				t.Error("A panic was expected when call Covariance with empty parameters")
			}
		}()

		Covariance(c.x, c.y)
	}
}

func TestCorrelation(t *testing.T) {
	cases := []struct {
		x, y collections.Vector
		want float64
	}{
		{
			collections.Vector{1.0, 2.0, 3.0},
			collections.Vector{1.0, 2.0},
			0.35355339059327373,
		},
		{
			collections.Vector{1.0, 2.0},
			collections.Vector{1.0, 2.0},
			0.9999999999999999,
		},
		{
			collections.Vector{1.0, 2.0, 3.0},
			collections.Vector{1.0, 2.0, 3.0},
			1.0,
		},
		{
			collections.Vector{1.0, 1.0},
			collections.Vector{0.0, 0.0},
			0.0,
		},
	}
	for _, c := range cases {
		got := Correlation(c.x, c.y)
		if got != c.want {
			t.Errorf("Correlation(%v, %v) want: %v but got: %v",
				c.x, c.y, c.want, got)
		}
	}
}

func TestCorrelation_WhenVectorHasOneElement(t *testing.T) {
	cases := []struct {
		x, y collections.Vector
	}{
		{collections.Vector{1.0}, collections.Vector{1.0, 2.0, 3.0}},
		{collections.Vector{}, collections.Vector{1.0, 2.0, 3.0}},
	}

	for _, c := range cases {
		defer func() {
			if recover() == nil {
				t.Error("A panic was expected when call Covariance with empty parameters")
			}
		}()

		Correlation(c.x, c.y)
	}
}
