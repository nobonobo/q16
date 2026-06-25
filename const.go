package q16

const (
	Zero Fixed = 0

	// Scale Q16.16 のスケールファクター (2^16)
	Scale     = 1 << 16
	ShiftBits = 16

	// MaxInt32 / MinInt32 int32 の範囲 (std math package と名前衝突するため定義)
	MaxInt32 = 1<<31 - 1
	MinInt32 = -(1 << 31)
)

var (
	// MaxFixed / MinFixed Q16.16 の範囲
	MaxFixed Fixed = Fixed(MaxInt32)
	MinFixed Fixed = Fixed(MinInt32)

	// Pi π ≈ 3.141592653589793
	Pi     = FromFloat64(3.14159265358979323846)
	Period = Mul(FromInt(2), Pi) // 2π
)
