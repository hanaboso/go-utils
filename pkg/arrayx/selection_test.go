package arrayx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{
			arr:      []int{11, 22, 33},
			expected: 22,
		},
		{
			arr:      []int{11},
			expected: 1,
		},
		{
			arr:      []int{},
			expected: 1,
		},
		{
			arr:      nil,
			expected: 1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Get(test.arr, 1, 1))
	}
}
