package math

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Abs[T numeric](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

func Sum[T numeric](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}

	return sum
}
