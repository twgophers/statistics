package statistics

import (
	"testing"
)

func TestMaxIn(t *testing.T) {
	cases := []struct {
		sample []float64
		wanted float64
	}{
		{
			[]float64{11.0},
			11.0,
		},
		{
			[]float64{11.0, 12.0},
			12.0,
		},
		{
			[]float64{11.0, 13.0, 12.0},
			13.0,
		},
	}
	for _, c := range cases {
		gotMax := MaxIn(c.sample)
		if gotMax != c.wanted {
			t.Errorf("Expected max (%v) in (%v) but got (%v)", c.wanted, c.sample, gotMax)
		}
	}
}

func TestMaxInFail(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected panic when empty sample")
		}
	}()

	MaxIn([]float64{})
}

func TestMinIn(t *testing.T) {
	cases := []struct {
		sample []float64
		wanted float64
	}{
		{
			[]float64{13.0},
			13.0,
		},
		{
			[]float64{12.0, 13.0},
			12.0,
		},
		{
			[]float64{12.0, 11.0, 13.0},
			11.0,
		},
	}
	for _, c := range cases {
		gotMin := MinIn(c.sample)
		if gotMin != c.wanted {
			t.Errorf("Expected max (%v) in (%v) but got (%v)", c.wanted, c.sample, gotMin)
		}
	}
}

func TestMinInFail(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected panic when empty sample")
		}
	}()

	MinIn([]float64{})
}

func TestSum(t *testing.T) {
	cases := []struct {
		sample []float64
		wanted float64
	}{
		{
			[]float64{7.0},
			7.0,
		},
		{
			[]float64{32.0, 7.0},
			39.0,
		},
		{
			[]float64{},
			0.0,
		},
	}
	for _, c := range cases {
		gotSum := Sum(c.sample)
		if gotSum != c.wanted {
			t.Errorf("Expected total (%v) summing up (%v) but got (%v)", c.wanted, c.sample, gotSum)
		}
	}
}
