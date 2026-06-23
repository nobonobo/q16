package q16

// Mask Q16.16 の小数部マスク (下位16ビット)
const Mask = Fixed((1 << ShiftBits) - 1)

// Floor xの床関数 (小数点以下切り捨て、負方向に丸め)
func Floor(x Fixed) Fixed {
	return (x >> ShiftBits) << ShiftBits
}

// Ceil xの天井関数 (小数点以下切り上げ)
func Ceil(x Fixed) Fixed {
	if x&Mask == 0 {
		return x
	}
	return Floor(x) + Scale
}

// Round xを四捨五入 (最も近い整数へ、.5 は正方向に丸め)
func Round(x Fixed) Fixed {
	return Floor(x + Fixed(Scale/2))
}

// Trunc xの切り捨て
func Trunc(x Fixed) Fixed {
	if x >= 0 {
		return x &^ Mask
	}
	return -((-x) &^ Mask)
}
