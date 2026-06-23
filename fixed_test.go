package q16_test

import (
	"testing"
	"time"

	"github.com/nobonobo/q16"
)

func TestFromDuration(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		d    time.Duration
		want q16.Fixed
	}{
		// test cases
		{"1 second", time.Second, q16.FromInt(1)},
		{"100ms", 100 * time.Millisecond, q16.FromFloat32(0.1)},
		{"500ms", 500 * time.Millisecond, q16.FromFloat32(0.5)},
		{"1.5s", 1500 * time.Millisecond, q16.FromFloat32(1.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q16.FromDuration(tt.d)
			if got != tt.want {
				t.Errorf("FromDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
