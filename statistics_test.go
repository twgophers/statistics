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
