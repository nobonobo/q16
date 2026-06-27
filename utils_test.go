package q16_test

import (
	"testing"

	"github.com/nobonobo/q16"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		x        q16.Fixed
		inMin    q16.Fixed
		inMax    q16.Fixed
		outMin   q16.Fixed
		outMax   q16.Fixed
		expected q16.Fixed
	}{
		{
			name:     "normal conversion 0-100 to 0-10",
			x:        q16.FromInt(50),
			inMin:    q16.FromInt(0),
			inMax:    q16.FromInt(100),
			outMin:   q16.FromInt(0),
			outMax:   q16.FromInt(10),
			expected: q16.FromFloat32(5.0),
		},
		{
			name:     "convert 25 in range 0-100 to 0-10",
			x:        q16.FromInt(25),
			inMin:    q16.FromInt(0),
			inMax:    q16.FromInt(100),
			outMin:   q16.FromInt(0),
			outMax:   q16.FromInt(10),
			expected: q16.FromFloat32(2.5),
		},
		{
			name:     "convert 75 in range 0-100 to 0-10",
			x:        q16.FromInt(75),
			inMin:    q16.FromInt(0),
			inMax:    q16.FromInt(100),
			outMin:   q16.FromInt(0),
			outMax:   q16.FromInt(10),
			expected: q16.FromFloat32(7.5),
		},
		{
			name:     "inMin equals inMax returns outMin",
			x:        q16.FromInt(50),
			inMin:    q16.FromInt(10),
			inMax:    q16.FromInt(10),
			outMin:   q16.FromInt(0),
			outMax:   q16.FromInt(10),
			expected: q16.FromInt(0),
		},
		{
			name:     "negative range conversion",
			x:        q16.FromInt(-50),
			inMin:    q16.FromInt(-100),
			inMax:    q16.FromInt(100),
			outMin:   q16.FromInt(-10),
			outMax:   q16.FromInt(10),
			expected: q16.FromInt(-5),
		},
		{
			name:     "reverse conversion 100-0 to 0-10",
			x:        q16.FromInt(50),
			inMin:    q16.FromInt(100),
			inMax:    q16.FromInt(0),
			outMin:   q16.FromInt(0),
			outMax:   q16.FromInt(10),
			expected: q16.FromFloat32(5.0),
		},
		{
			name:     "convert first value",
			x:        q16.FromInt(0),
			inMin:    q16.FromInt(0),
			inMax:    q16.FromInt(100),
			outMin:   q16.FromInt(-5),
			outMax:   q16.FromInt(5),
			expected: q16.FromInt(-5),
		},
		{
			name:     "convert last value",
			x:        q16.FromInt(100),
			inMin:    q16.FromInt(0),
			inMax:    q16.FromInt(100),
			outMin:   q16.FromInt(-5),
			outMax:   q16.FromInt(5),
			expected: q16.FromInt(5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := q16.Map(tt.x, tt.inMin, tt.inMax, tt.outMin, tt.outMax)
			if result != tt.expected {
				t.Errorf("Map(%v, %v, %v, %v, %v) = %v, want %v",
					tt.x, tt.inMin, tt.inMax, tt.outMin, tt.outMax, result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		initial  q16.Fixed
		args     []q16.Fixed
		expected q16.Fixed
	}{
		{
			name:     "find max in multiple values",
			initial:  q16.FromInt(1),
			args:     []q16.Fixed{q16.FromInt(5), q16.FromInt(3), q16.FromInt(9), q16.FromInt(2)},
			expected: q16.FromInt(9),
		},
		{
			name:     "no args returns initial",
			initial:  q16.FromInt(5),
			args:     []q16.Fixed{},
			expected: q16.FromInt(5),
		},
		{
			name:     "initial is the max",
			initial:  q16.FromInt(10),
			args:     []q16.Fixed{q16.FromInt(3), q16.FromInt(7), q16.FromInt(2)},
			expected: q16.FromInt(10),
		},
		{
			name:     "last arg is the max",
			initial:  q16.FromInt(1),
			args:     []q16.Fixed{q16.FromInt(2), q16.FromInt(3), q16.FromInt(5)},
			expected: q16.FromInt(5),
		},
		{
			name:     "negative values",
			initial:  q16.FromInt(-10),
			args:     []q16.Fixed{q16.FromInt(-3), q16.FromInt(-7), q16.FromInt(-1)},
			expected: q16.FromInt(-1),
		},
		{
			name:     "all same values",
			initial:  q16.FromInt(5),
			args:     []q16.Fixed{q16.FromInt(5), q16.FromInt(5)},
			expected: q16.FromInt(5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := q16.Max(tt.initial, tt.args...)
			if result != tt.expected {
				t.Errorf("Max(%v, %v) = %v, want %v", tt.initial, tt.args, result, tt.expected)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		initial  q16.Fixed
		args     []q16.Fixed
		expected q16.Fixed
	}{
		{
			name:     "find min in multiple values",
			initial:  q16.FromInt(5),
			args:     []q16.Fixed{q16.FromInt(1), q16.FromInt(3), q16.FromInt(9), q16.FromInt(2)},
			expected: q16.FromInt(1),
		},
		{
			name:     "no args returns initial",
			initial:  q16.FromInt(5),
			args:     []q16.Fixed{},
			expected: q16.FromInt(5),
		},
		{
			name:     "initial is the min",
			initial:  q16.FromInt(-10),
			args:     []q16.Fixed{q16.FromInt(-3), q16.FromInt(-7), q16.FromInt(-2)},
			expected: q16.FromInt(-10),
		},
		{
			name:     "last arg is the min",
			initial:  q16.FromInt(5),
			args:     []q16.Fixed{q16.FromInt(3), q16.FromInt(2), q16.FromInt(1)},
			expected: q16.FromInt(1),
		},
		{
			name:     "negative values",
			initial:  q16.FromInt(-1),
			args:     []q16.Fixed{q16.FromInt(-10), q16.FromInt(-3), q16.FromInt(-7)},
			expected: q16.FromInt(-10),
		},
		{
			name:     "all same values",
			initial:  q16.FromInt(5),
			args:     []q16.Fixed{q16.FromInt(5), q16.FromInt(5)},
			expected: q16.FromInt(5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := q16.Min(tt.initial, tt.args...)
			if result != tt.expected {
				t.Errorf("Min(%v, %v) = %v, want %v", tt.initial, tt.args, result, tt.expected)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name     string
		x        q16.Fixed
		min      q16.Fixed
		max      q16.Fixed
		expected q16.Fixed
	}{
		{
			name:     "value in range returns as is",
			x:        q16.FromInt(5),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(5),
		},
		{
			name:     "value at min boundary",
			x:        q16.FromInt(0),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(0),
		},
		{
			name:     "value at max boundary",
			x:        q16.FromInt(10),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(0),
		},
		{
			name:     "value above max wraps around",
			x:        q16.FromInt(12),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(2),
		},
		{
			name:     "value below min wraps around",
			x:        q16.FromInt(-12),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(8),
		},
		{
			name:     "value far above max wraps around",
			x:        q16.FromInt(25),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(5),
		},
		{
			name:     "value far below min wraps around",
			x:        q16.FromInt(-15),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(5),
		},
		{
			name:     "max <= min returns x",
			x:        q16.FromInt(5),
			min:      q16.FromInt(10),
			max:      q16.FromInt(5),
			expected: q16.FromInt(5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := q16.Wrap(tt.x, tt.min, tt.max)
			if result != tt.expected {
				t.Errorf("Wrap(%v, %v, %v) = %v, want %v", tt.x, tt.min, tt.max, result, tt.expected)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		name     string
		x        q16.Fixed
		min      q16.Fixed
		max      q16.Fixed
		expected q16.Fixed
	}{
		{
			name:     "value in range returns as is",
			x:        q16.FromInt(5),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(5),
		},
		{
			name:     "value at min boundary",
			x:        q16.FromInt(0),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(0),
		},
		{
			name:     "value at max boundary",
			x:        q16.FromInt(10),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(10),
		},
		{
			name:     "value below min clamps to min",
			x:        q16.FromInt(-5),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(0),
		},
		{
			name:     "value above max clamps to max",
			x:        q16.FromInt(15),
			min:      q16.FromInt(0),
			max:      q16.FromInt(10),
			expected: q16.FromInt(10),
		},
		{
			name:     "far below min clamps to min",
			x:        q16.FromInt(-100),
			min:      q16.FromInt(-50),
			max:      q16.FromInt(50),
			expected: q16.FromInt(-50),
		},
		{
			name:     "far above max clamps to max",
			x:        q16.FromInt(100),
			min:      q16.FromInt(-50),
			max:      q16.FromInt(50),
			expected: q16.FromInt(50),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := q16.Clamp(tt.x, tt.min, tt.max)
			if result != tt.expected {
				t.Errorf("Clamp(%v, %v, %v) = %v, want %v", tt.x, tt.min, tt.max, result, tt.expected)
			}
		})
	}
}
