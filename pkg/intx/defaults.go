package intx

import (
	"github.com/hanaboso/go-utils/pkg/mathx"
	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Float | constraints.Integer
}

func Default[T Numeric](value, def T) T {
	return mathx.Default(value, def)
}
