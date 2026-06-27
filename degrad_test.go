package q16

import (
	"fmt"
	"testing"
)

func TestDegToRadAccuracy(t *testing.T) {
	// 30度 = π/6 ≈ 0.5235987756 radian
	// Q16.16では: 0.5235987756 * 65536 ≈ 34279 (0x85C7)
	result30 := DegToRad(FromInt(30))
	fmt.Printf("DegToRad(30) = %d (float: %.6f), expected: 34279 (float: 0.523599)\n", result30, float64(result30)/Scale)

	// 90度 = π/2 ≈ 1.5707963268 radian
	// Q16.16では: 1.5707963268 * 65536 ≈ 102938 (0x1921F)
	result90 := DegToRad(FromInt(90))
	fmt.Printf("DegToRad(90) = %d (float: %.6f), expected: 102938 (float: 1.570796)\n", result90, float64(result90)/Scale)

	// π定数の確認
	pi := FromFloat32(3.14159265358979)
	fmt.Printf("Pi (Q16.16)  = %d (float: %.6f)\n", pi, float64(pi)/Scale)

	// π/2の確認
	piHalf := FromFloat32(1.57079632679490)
	fmt.Printf("Pi/2 (Q16.16) = %d (float: %.6f)\n", piHalf, float64(piHalf)/Scale)

	// 90度の結果がπ/2と等しいか確認
	if result90 != piHalf {
		t.Errorf("DegToRad(90) = %d, want π/2 = %d", result90, piHalf)
	}
}
