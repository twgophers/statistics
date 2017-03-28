package statistics

import (
	"reflect"
	"testing"
)

func TestValues(t *testing.T) {
	cases := []struct {
		sample []float64
		wanted []int
	}{
		{[]float64{1.0, 1.0}, []int{2}},
		{[]float64{1.0, 2.0}, []int{1, 1}},
		{[]float64{}, []int{}},
	}
	for _, c := range cases {
		got := NewCounter(c.sample).Values()
		if !reflect.DeepEqual(got, c.wanted) {
			t.Errorf("Counter.values(%v) want: %v; got %v", c.sample, c.wanted, got)
		}
	}
}
