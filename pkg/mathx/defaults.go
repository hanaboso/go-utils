package mathx

import "golang.org/x/exp/constraints"

type Numeric interface {
	constraints.Float | constraints.Integer
}

func Default[T Numeric](value, def T) T {
	if value == 0 {
		return def
	}

	return value
}
