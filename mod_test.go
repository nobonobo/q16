package q16

import "testing"

func TestMod(t *testing.T) {
	tests := []struct {
		name         string
		a, b         Fixed
		expectQuot   float32
		expectRemain float32
	}{
		{"positive: a=5, b=2", FromFloat32(5), FromFloat32(2), 2, 1},
		{"negative: a=-5, b=2", FromFloat32(-5), FromFloat32(2), -2, -1},
		{"positive: a=0.9π, b=π", Mul(Pi, FromFloat32(1.9)), Pi, 1, 0.9 * float32(Pi.Float32())},
		{"negative: a=-0.5π, b=π", Mul(Pi, FromFloat32(-2.5)), Pi, -2, -0.5 * float32(Pi.Float32())},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			q, r := DivMod(tc.a, tc.b)
			qGot := q.Float32()
			rGot := r.Float32()
			if abs(qGot-tc.expectQuot) > 0.01 || abs(rGot-tc.expectRemain) > 0.01 {
				t.Errorf("Mod(%f, %f): got (%.4f, %.4f), want (%.4f, %.4f)", tc.a.Float32(), tc.b.Float32(), qGot, rGot, tc.expectQuot, tc.expectRemain)
			}
		})
	}
}

func abs(f float32) float32 {
	if f < 0 {
		return -f
	}
	return f
}
