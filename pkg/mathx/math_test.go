package mathx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		values []int
		result int
	}{
		{
			values: []int{1, 2, 3},
			result: 3,
		},
		{
			values: []int{11, 2},
			result: 11,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, Max(test.values[0], test.values[1:]...))
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		values []int
		result int
	}{
		{
			values: []int{1, 2, 3},
			result: 1,
		},
		{
			values: []int{11, 2},
			result: 2,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, Min(test.values[0], test.values[1:]...))
	}
}
