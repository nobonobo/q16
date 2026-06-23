package q16_test

import (
	"testing"

	"github.com/nobonobo/q16"
)

func TestFloor(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    q16.Fixed
		want q16.Fixed
	}{
		// TODO: Add test cases.
		{"test case 1", q16.FromFloat32(3.7), q16.FromInt(3)},
		{"test case 2", q16.FromFloat32(-3.7), q16.FromInt(-4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Floor(tt.x)
			if got != tt.want {
				t.Errorf("Floor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    q16.Fixed
		want q16.Fixed
	}{
		// TODO: Add test cases.
		{"test case 1", q16.FromFloat32(3.7), q16.FromInt(4)},
		{"test case 2", q16.FromFloat32(-3.7), q16.FromInt(-3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Ceil(tt.x)
			if got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    q16.Fixed
		want q16.Fixed
	}{
		{"test case 1", q16.FromFloat32(3.7), q16.FromInt(4)},
		{"test case 2", q16.FromFloat32(-3.7), q16.FromInt(-4)},
		{"test case 3", q16.FromFloat32(3.4), q16.FromInt(3)},
		{"test case 4", q16.FromFloat32(-3.4), q16.FromInt(-3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Round(tt.x)
			if got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrunc(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    q16.Fixed
		want q16.Fixed
	}{
		{"test case 1", q16.FromFloat32(3.7), q16.FromInt(3)},
		{"test case 2", q16.FromFloat32(-3.7), q16.FromInt(-3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Trunc(tt.x)
			if got != tt.want {
				t.Errorf("Trunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
