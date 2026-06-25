package q16

// Q16.16 定数 (std math 依存なし)
var (
	pi         = FromFloat64(3.14159265358979323846)
	pi2        = FromFloat64(6.28318530717958647693) // 2π
	piHalf     = FromFloat64(1.57079632679489661923) // π/2
	piQuart    = FromFloat64(0.78539816339744830962) // π/4
	negPiQuart = FromFloat64(-0.78539816339744830962)
)

// Sin sin(x) を計算する (ラジアン入力 Q16.16 → Q16.16 出力 [-1, 1])
func Sin(x Fixed) Fixed {
	if x == 0 {
		return 0
	}
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}

	x = x % pi2

	if x >= pi {
		neg = !neg
		x = x - pi
	}

	if x > piHalf {
		x = pi - x
	}

	var result Fixed
	if x > piQuart {
		result = cosSmall(piHalf - x)
	} else {
		result = sinSmall(x)
	}

	if neg {
		return -result
	}
	return result
}

// Cos cos(x) を計算する (ラジアン入力 Q16.16 → Q16.16 出力 [-1, 1])
func Cos(x Fixed) Fixed {
	if x == 0 {
		return FromInt(1)
	}
	x = Abs(x)
	x = x % pi2

	neg := false
	if x >= pi {
		neg = !neg
		x = x - pi
	}

	if x > piHalf {
		neg = !neg
		x = pi - x
	}

	var result Fixed
	if x > piQuart {
		result = sinSmall(piHalf - x)
	} else {
		result = cosSmall(x)
	}

	if neg {
		return -result
	}
	return result
}

// Tan tan(x) を計算する (ラジアン入力 Q16.16 → Q16.16)
func Tan(x Fixed) Fixed {
	return divRound(Sin(x), Cos(x))
}

// Atan atan(x) を計算する (Q16.16 → Q16.16 ラジアン出力)
func Atan(x Fixed) Fixed {
	negate := false
	if x < 0 {
		x = -x
		negate = true
	}

	// |x| > 1 の場合: atan(x) = π/2 - atan(1/x)
	if x > FromInt(1) {
		recip := divRound(FromInt(1), x)
		result := atanSmall(recip)
		if negate {
			return -(piHalf - result)
		}
		return piHalf - result
	}

	result := atanSmall(x)
	if negate {
		return -result
	}
	return result
}

// --- 内部関数 ---

// sinSmall |x| <= π/4 の範囲で sin を計算 (テイラー展開)
func sinSmall(x Fixed) Fixed {
	x2 := mulRound(x, x)
	result := Zero
	term := x

	for i := 1; i <= 9; i += 2 {
		result = result + term
		denom := FromInt((i + 1) * (i + 2))
		term = mulRound(term, x2)
		term = divRound(-term, denom)
	}

	return result
}

// cosSmall |x| <= π/4 の範囲で cos を計算 (テイラー展開)
func cosSmall(x Fixed) Fixed {
	x2 := mulRound(x, x)
	result := Zero
	term := FromInt(1) // 最初の項 = 1

	for k := 1; k <= 5; k++ {
		result = result + term
		denom := FromInt((2 * k) * (2*k - 1))
		term = mulRound(term, x2)
		term = divRound(-term, denom)
	}

	return result
}

// atanSmall |x| <= 1 の範囲で atan を計算 (多項式近似)
func atanSmall(x Fixed) Fixed {
	x2 := mulRound(x, x)

	// atan(x) ≈ x * (π/4 + a*x²) / (1 + b*x²)
	// a = b * π/4 となるように設計し、x=1 で完全に π/4 に一致させる
	a := FromFloat32(0.27269)
	b := FromFloat32(0.3472)

	num := piQuart + mulRound(a, x2)
	den := FromInt(1) + mulRound(b, x2)

	return mulRound(x, divRound(num, den))
}
