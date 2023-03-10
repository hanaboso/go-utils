package intx

import (
	"github.com/hanaboso/go-utils/pkg/mathx"
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a T, b ...T) T {
	return mathx.Max(a, b...)
}

func Min[T constraints.Ordered](a T, b ...T) T {
	return mathx.Min(a, b...)
}
