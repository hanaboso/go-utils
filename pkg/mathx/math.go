package mathx

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a T, b ...T) T {
	max := a
	for _, i := range b {
		if i > max {
			max = i
		}
	}

	return max
}

func Min[T constraints.Ordered](a T, b ...T) T {
	min := a
	for _, i := range b {
		if i < min {
			min = i
		}
	}

	return min
}
