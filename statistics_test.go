package statistics

import (
	"testing"
)

type SampleTestTable struct {
	sample Sample
	wanted float64
}

func TestMax(t *testing.T) {
	cases := []SampleTestTable{
		{
			Sample{11.0},
			11.0,
		},
		{
			Sample{11.0, 12.0},
			12.0,
		}, {
			Sample{11.0, 13.0, 12.0},
			13.0,
		},
	}
	for _, c := range cases {
		gotMax := c.sample.Max()

		if gotMax != c.wanted {
			t.Errorf("Expected max (%v) in (%v) but got (%v)", c.wanted, c.sample, gotMax)
		}
	}
}

func TestMaxFail(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected panic when empty sample")
		}
	}()

	Sample{}.Max()
}

func TestMin(t *testing.T) {
	cases := []SampleTestTable{
		{
			Sample{13.0},
			13.0,
		},
		{
			Sample{12.0, 13.0},
			12.0,
		},
		{
			Sample{12.0, 11.0, 13.0},
			11.0,
		},
	}
	for _, c := range cases {
		gotMin := c.sample.Min()
		if gotMin != c.wanted {
			t.Errorf("Expected max (%v) in (%v) but got (%v)", c.wanted, c.sample, gotMin)
		}
	}
}

func TestMinFail(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected panic when empty sample")
		}
	}()

	Sample{}.Min()
}

func TestSum(t *testing.T) {
	cases := []SampleTestTable{
		{
			Sample{7.0},
			7.0,
		},
		{
			Sample{32.0, 7.0},
			39.0,
		},
		{
			Sample{},
			0.0,
		},
	}
	for _, c := range cases {
		gotSum := c.sample.Sum()
		if gotSum != c.wanted {
			t.Errorf("Expected total (%v) summing up (%v) but got (%v)", c.wanted, c.sample, gotSum)
		}
	}
}

func TestMean(t *testing.T) {
	cases := []SampleTestTable{
		{
			Sample{7.0},
			7.0,
		},
		{
			Sample{13.0, 14.0},
			13.5,
		},
	}
	for _, c := range cases {
		gotMean := c.sample.Mean()

		if gotMean != c.wanted {
			t.Errorf("Expected mean of (%v) for (%v) but got (%v)", c.wanted, c.sample, gotMean)
		}
	}
}

func TestMeanPanicsWhenEmptySlice(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected mean panic when empty sample")
		}
	}()

	Sample{}.Mean()
}

func TestMedian(t *testing.T) {
	cases := []SampleTestTable{
		{
			Sample{7.0},
			7.0,
		},
		{
			Sample{8.0, 11.0},
			9.5,
		},
		{
			Sample{8.0, 11.0, 7.0},
			11.0,
		},
		{
			Sample{8.0, 10.0, 11.0, 7.0},
			10.5,
		},
	}
	for _, c := range cases {
		gotMedian := c.sample.Median()

		if gotMedian != c.wanted {
			t.Errorf("Expected median (%v) for (%v) but got (%v)", c.wanted, c.sample, gotMedian)
		}
	}
}

func TestMedianPanicsWhenEmptySlice(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected median panic when empty sample")
		}
	}()

	Sample{}.Median()
}
