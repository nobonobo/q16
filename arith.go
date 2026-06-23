package q16

// Mul 2つのFixedを乗算する (Q16.16 * Q16.16 → Q16.16)
// int64で計算してオーバーフロー防止
func Mul(a, b Fixed) Fixed {
	prod := int64(a) * int64(b) // 最大 ~2^30 * 2^30 = 2^60 (int64に収まる)
	return Fixed(prod >> ShiftBits)
}

// Div 2つのFixedで除算する (Q16.16 / Q16.16 → Q16.16)
func Div(a, b Fixed) Fixed {
	// int64で分子を拡張してから除算
	num := int64(a) << ShiftBits // 最大 ~2^30 * 2^16 = 2^46 (int64に収まる)
	return Fixed(num / int64(b))
}

// Sign 固定小数点数の符号を取得する
func Sign(a Fixed) Fixed {
	if a > Zero {
		return FromInt(1)
	} else if a < Zero {
		return FromInt(-1)
	} else {
		return FromInt(0)
	}
}
