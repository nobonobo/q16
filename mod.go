package q16

// DivMod 2つのFixedの商(quotient)と余り(remainder)を計算する (Q16.16)
// Divと同様の除算（ゼロ方向への切り捨て）を行い、剰余も同時に返す。
// 符号規則: remainder の符号は a に従い、|remainder| < |b| を満たす。
func DivMod(a, b Fixed) (quotient, remainder Fixed) {
	q := Trunc(Div(a, b))
	// r = a - q * b （q16.Mul で乗算）
	r := a - Mul(q, b)
	return q, r
}
