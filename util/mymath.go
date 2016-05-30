package mymath

const MaxInt64 = (1 << (63-1))
const MinInt64 = -(1 << (63-1))

func Min(x, y int64) int64 {

	if x < y {
		return x
	}
	return y
}

func Max(x, y int64) int64 {

	if x > y {
		return x
	}
	return y
}