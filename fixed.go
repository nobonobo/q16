package q16

import "time"

// Fixed Q16.16 固定小数点数
type Fixed int32

// DegToRad degreeからradianに変換する (固定小数点数のみで高精度に計算)
func DegToRad(deg Fixed) Fixed {
	// π/180 ≈ 0.0174532925199433 を 2^30 スケールで表すと 18740263
	// int64の乗算を行い、四捨五入して 30ビット右シフトする
	prod := int64(deg) * 18740263
	if prod >= 0 {
		return Fixed((prod + (1 << 29)) >> 30)
	}
	return Fixed((prod - (1 << 29)) >> 30)
}

// RadToDeg radianからdegreeに変換する (固定小数点数のみで高精度に計算)
func RadToDeg(rad Fixed) Fixed {
	// 180/π ≈ 57.29577951308232 を 2^24 スケールで表すと 955505
	// int64の乗算を行い、四捨五入して 24ビット右シフトする
	prod := int64(rad) * 955505
	if prod >= 0 {
		return Fixed((prod + (1 << 23)) >> 24)
	}
	return Fixed((prod - (1 << 23)) >> 24)
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
