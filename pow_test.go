package q16_test

import (
	"testing"

	"github.com/nobonobo/q16"
)

func TestPow(t *testing.T) {
	tests := []struct {
		name string
		base q16.Fixed
		exp  q16.Fixed
		want q16.Fixed
	}{
		{"1^1 = 1", q16.FromInt(1), q16.FromInt(1), q16.FromInt(1)},
		{"2^1 = 2", q16.FromInt(2), q16.FromInt(1), q16.FromInt(2)},
		{"2^2 = 4", q16.FromInt(2), q16.FromInt(2), q16.FromInt(4)},
		{"5^0 = 1", q16.FromInt(5), q16.FromInt(0), q16.FromInt(1)},
		{"1^5 = 1", q16.FromInt(1), q16.FromInt(5), q16.FromInt(1)},
		{"4^0.5 = 2", q16.FromInt(4), q16.FromFloat32(0.5), q16.FromInt(2)},
		{"2^-1 = 0.5", q16.FromInt(2), q16.FromInt(-1), q16.FromFloat32(0.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Pow(tt.base, tt.exp)
			// 1 LSB 程度の僅かな丸め誤差を考慮し、1 LSB 以内の差分を許容
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if diff > 1 {
				t.Errorf("Pow() = %v (float: %f), want %v (float: %f)", got, got.Float32(), tt.want, tt.want.Float32())
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		name string
		x    q16.Fixed
		want q16.Fixed
	}{
		{"sqrt(0) = 0", q16.FromInt(0), q16.FromInt(0)},
		{"sqrt(1) = 1", q16.FromInt(1), q16.FromInt(1)},
		{"sqrt(4) = 2", q16.FromInt(4), q16.FromInt(2)},
		{"sqrt(9) = 3", q16.FromInt(9), q16.FromInt(3)},
		{"sqrt(16) = 4", q16.FromInt(16), q16.FromInt(4)},
		{"sqrt(25) = 5", q16.FromInt(25), q16.FromInt(5)},
		{"sqrt(36) = 6", q16.FromInt(36), q16.FromInt(6)},
		{"sqrt(49) = 7", q16.FromInt(49), q16.FromInt(7)},
		{"sqrt(64) = 8", q16.FromInt(64), q16.FromInt(8)},
		{"sqrt(81) = 9", q16.FromInt(81), q16.FromInt(9)},
		{"sqrt(100) = 10", q16.FromInt(100), q16.FromInt(10)},
		{"sqrt(0.25) = 0.5", q16.FromFloat32(0.25), q16.FromFloat32(0.5)},
		{"sqrt(-4) = 0", q16.FromInt(-4), q16.Zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Sqrt(tt.x)
			if got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExp(t *testing.T) {
	tests := []struct {
		name string
		x    q16.Fixed
		want q16.Fixed
	}{
		{"e^0 = 1", q16.FromInt(0), q16.FromInt(1)},
		{"e^1 = e", q16.FromInt(1), q16.E},
		{"e^-1 = 1/e", q16.FromInt(-1), q16.Div(q16.FromInt(1), q16.E)},
		{"e^11 (overflow limit) = MaxFixed", q16.FromInt(11), q16.MaxFixed},
		{"e^-12 (underflow limit) = 0", q16.FromInt(-12), q16.Zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Exp(tt.x)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if diff > 1 {
				t.Errorf("Exp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog(t *testing.T) {
	tests := []struct {
		name string
		x    q16.Fixed
		want q16.Fixed
	}{
		{"ln(1) = 0", q16.FromInt(1), q16.FromInt(0)},
		{"ln(e) = 1", q16.E, q16.FromInt(1)},
		{"ln(0) = MinFixed", q16.FromInt(0), q16.MinFixed},
		{"ln(-5) = MinFixed", q16.FromInt(-5), q16.MinFixed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Log(tt.x)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if diff > 1 {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog2(t *testing.T) {
	tests := []struct {
		name string
		x    q16.Fixed
		want q16.Fixed
	}{
		{"log2(1) = 0", q16.FromInt(1), q16.FromInt(0)},
		{"log2(2) = 1", q16.FromInt(2), q16.FromInt(1)},
		{"log2(4) = 2", q16.FromInt(4), q16.FromInt(2)},
		{"log2(0.5) = -1", q16.FromFloat32(0.5), q16.FromInt(-1)},
		{"log2(0) = MinFixed", q16.FromInt(0), q16.MinFixed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Log2(tt.x)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if diff > 1 {
				t.Errorf("Log2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog10(t *testing.T) {
	tests := []struct {
		name string
		x    q16.Fixed
		want q16.Fixed
	}{
		{"log10(1) = 0", q16.FromInt(1), q16.FromInt(0)},
		{"log10(10) = 1", q16.FromInt(10), q16.FromInt(1)},
		{"log10(100) = 2", q16.FromInt(100), q16.FromInt(2)},
		{"log10(0) = MinFixed", q16.FromInt(0), q16.MinFixed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Log10(tt.x)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if diff > 1 {
				t.Errorf("Log10() = %v, want %v", got, tt.want)
			}
		})
	}
}
