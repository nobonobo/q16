package q16

import "time"

// Fixed Q16.16 固定小数点数
type Fixed int32

// DegToRad degreeからradianに変換する (固定小数点数のみで高精度に計算)
func DegToRad(deg Fixed) Fixed {
	// deg × π / 180 を計算
	// deg × π → Q32.32 → >> 16 で Q16.16
	prod := int64(deg) * int64(pi)
	result := Fixed((prod + (1 << 15)) >> 16) // deg × π in Q16.16
	return divRound(result, FromInt(180))     // / 180
}

// RadToDeg radianからdegreeに変換する (固定小数点数のみで高精度に計算)
func RadToDeg(rad Fixed) Fixed {
	// rad × 180 / π を計算
	prod := int64(rad) * 180
	return divRound(Fixed(prod), pi)
}

// FromInt int から Fixed を作成する
func FromInt(i int) Fixed {
	return Fixed(i << ShiftBits)
}

// FromFloat32 float32 から Fixed を作成する (四捨五入)
func FromFloat32(f float32) Fixed {
	v := int64(float64(f)*Scale + 0.5)
	if v > MaxInt32 {
		return MaxFixed
	}
	if v < MinInt32 {
		return MinFixed
	}
	return Fixed(v)
}

func FromDuration(d time.Duration) Fixed {
	// Convert duration (nanoseconds) to Fixed (seconds) without using float.
	// time.Second = 1_000_000_000 ns.
	// Separate whole seconds and nanosecond remainder to avoid overflow.
	secs := int64(d / time.Second) // truncate toward zero
	rem := int64(d % time.Second)  // remainder, same sign as d

	// Fixed representation of whole seconds.
	result := FromInt(int(secs))

	// Convert the remainder to Fixed with rounding.
	// Scale factor is 2^16.
	const halfNs int64 = 500_000_000 // half of 1e9 for rounding
	remFixed := rem * int64(Scale)
	if rem >= 0 {
		remFixed = (remFixed + halfNs) / 1_000_000_000
	} else {
		remFixed = (remFixed - halfNs) / 1_000_000_000
	}
	return result + Fixed(remFixed)
}

// FromFloat64 float64 から Fixed を作成する (四捨五入)
func FromFloat64(f float64) Fixed {
	v := int64(f*Scale + 0.5)
	if v > MaxInt32 {
		return MaxFixed
	}
	if v < MinInt32 {
		return MinFixed
	}
	return Fixed(v)
}

// Float32 Fixed を float32 に変換する
func (f Fixed) Float32() float32 {
	return float32(f) / Scale
}

// Float64 Fixed を float64 に変換する
func (f Fixed) Float64() float64 {
	return float64(f) / Scale
}

// Int Fixed を int に変換する
func (f Fixed) Int() int {
	return int(f >> ShiftBits)
}

// Abs |x| を計算する
func Abs(x Fixed) Fixed {
	if x < 0 {
		return -x
	}
	return x
}
