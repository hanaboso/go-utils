package mathx

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a T, b ...T) T {
	maxVar := a
	for _, i := range b {
		if i > maxVar {
			maxVar = i
		}
	}

	return maxVar
}

func Min[T constraints.Ordered](a T, b ...T) T {
	minVar := a
	for _, i := range b {
		if i < minVar {
			minVar = i
		}
	}

	return minVar
}
