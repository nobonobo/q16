package q16

func Map(x, inMin, inMax, outMin, outMax Fixed) Fixed {
	inRange := int64(inMax) - int64(inMin)
	if inRange == 0 {
		return outMin
	}
	outRange := int64(outMax) - int64(outMin)
	result := (int64(x)-int64(inMin))*outRange/inRange +
		int64(outMin)
	return Fixed(result)
}

func Max(in Fixed, args ...Fixed) Fixed {
	for _, v := range args {
		if v > in {
			in = v
		}
	}
	return in
}

func Min(in Fixed, args ...Fixed) Fixed {
	for _, v := range args {
		if v < in {
			in = v
		}
	}
	return in
}

func Wrap(x, min, max Fixed) Fixed {
	r := max - min
	if r <= 0 {
		return x
	}
	q := (x - min) / r
	if q >= 0 {
		return min + (x-min)%r
	}
	return max - (min-x)%r
}

func Clamp(x, min, max Fixed) Fixed {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
