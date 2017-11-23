package mpinverse

import (
	"math"
	"testing"
)

func TestGetCutoff(t *testing.T) {
	tests := []struct {
		inS1 []float64
		want float64
	}{
		{
			[]float64{
				1.0, 2.0, 3.0,
			},
			1e-15 * 3.0,
		},
		{
			[]float64{
				0.1, 0.2, 0.3,
			},
			1e-15 * 0.3,
		},
		{
			[]float64{
				1.0, 1.0, 1.0,
			},
			1e-15 * 1.0,
		},
		{
			[]float64{
				-1.0, -2.0, -3.0,
			},
			1e-15 * -1.0,
		},
		{
			[]float64{
				0.0,
			},
			0.0,
		},
	}

	for _, test := range tests {
		got := getCutoff(test.inS1)
		if math.Dim(test.want, got) > 1e-15 {
			t.Fatalf("want %f, but %f, %f:", test.want, got)
		}
	}
}
