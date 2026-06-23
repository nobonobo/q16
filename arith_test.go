package q16_test

import (
	"testing"

	"github.com/nobonobo/q16"
)

func TestMul(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    q16.Fixed
		b    q16.Fixed
		want q16.Fixed
	}{
		{"test case 1", q16.FromFloat32(5.0), q16.FromFloat32(2.0), q16.FromFloat32(10.0)},
		{"test case 2", q16.FromFloat32(3.0), q16.FromFloat32(2.0), q16.FromFloat32(6.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Mul(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    q16.Fixed
		b    q16.Fixed
		want q16.Fixed
	}{
		{"test case 1", q16.FromFloat32(5.0), q16.FromFloat32(2.0), q16.FromFloat32(2.5)},
		{"test case 2", q16.FromFloat32(3.0), q16.FromFloat32(2.0), q16.FromFloat32(1.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.Div(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}
