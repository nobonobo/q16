package q16_test

import (
	"testing"

	"github.com/nobonobo/q16"
)

func TestSin(t *testing.T) {
	tests := []struct {
		name string
		x    q16.Fixed
		want q16.Fixed
	}{
		{"sin(0) = 0", q16.FromInt(0), q16.FromInt(0)},
		{"sin(30deg) = 0.5", q16.DegToRad(q16.FromInt(30)), q16.FromFloat32(0.5)},
		{"sin(90deg) = 1", q16.DegToRad(q16.FromInt(90)), q16.FromInt(1)},
		{"sin(180deg) = 0", q16.DegToRad(q16.FromInt(180)), q16.FromInt(0)},
		{"sin(-90deg) = -1", q16.DegToRad(q16.FromInt(-90)), q16.FromInt(-1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Sin(tt.x)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			// 許容誤差 2 LSB
			if diff > 2 {
				t.Errorf("Sin() = %v (float: %f), want %v (float: %f)", got, got.Float32(), tt.want, tt.want.Float32())
			}
		})
	}
}

func TestCos(t *testing.T) {
	tests := []struct {
		name string
		x    q16.Fixed
		want q16.Fixed
	}{
		{"cos(0) = 1", q16.FromInt(0), q16.FromInt(1)},
		{"cos(60deg) = 0.5", q16.DegToRad(q16.FromInt(60)), q16.FromFloat32(0.5)},
		{"cos(90deg) = 0", q16.DegToRad(q16.FromInt(90)), q16.FromInt(0)},
		{"cos(180deg) = -1", q16.DegToRad(q16.FromInt(180)), q16.FromInt(-1)},
		{"cos(-180deg) = -1", q16.DegToRad(q16.FromInt(-180)), q16.FromInt(-1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Cos(tt.x)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			// 許容誤差 2 LSB
			if diff > 2 {
				t.Errorf("Cos() = %v (float: %f), want %v (float: %f)", got, got.Float32(), tt.want, tt.want.Float32())
			}
		})
	}
}

func TestTan(t *testing.T) {
	tests := []struct {
		name string
		x    q16.Fixed
		want q16.Fixed
	}{
		{"tan(0) = 0", q16.FromInt(0), q16.FromInt(0)},
		{"tan(45deg) = 1", q16.DegToRad(q16.FromInt(45)), q16.FromInt(1)},
		{"tan(-45deg) = -1", q16.DegToRad(q16.FromInt(-45)), q16.FromInt(-1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sinVal := q16.Sin(tt.x)
			cosVal := q16.Cos(tt.x)
			got := q16.Tan(tt.x)
			t.Logf("x: %v, Sin(x): %v, Cos(x): %v, Tan(x): %v", tt.x, sinVal, cosVal, got)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			// 許容誤差 5 LSB
			if diff > 5 {
				t.Errorf("Tan() = %v (float: %f), want %v (float: %f)", got, got.Float32(), tt.want, tt.want.Float32())
			}
		})
	}
}

func TestAtan(t *testing.T) {
	tests := []struct {
		name string
		x    q16.Fixed
		want q16.Fixed
	}{
		{"atan(0) = 0", q16.FromInt(0), q16.FromInt(0)},
		{"atan(1) = 45deg", q16.FromInt(1), q16.DegToRad(q16.FromInt(45))},
		{"atan(-1) = -45deg", q16.FromInt(-1), q16.DegToRad(q16.FromInt(-45))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Atan(tt.x)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			// 許容誤差 5 LSB
			if diff > 5 {
				t.Errorf("Atan() = %v (float: %f), want %v (float: %f)", got, got.Float32(), tt.want, tt.want.Float32())
			}
		})
	}
}
