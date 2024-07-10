package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisibleSumPairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		ar       []int32
		k        int32
		expected int32
	}{
		{
			name:     "test 1",
			n:        6,
			ar:       []int32{1, 2, 3, 4, 5, 6},
			k:        5,
			expected: 3,
		},
		{
			name:     "test 2",
			n:        6,
			ar:       []int32{1, 3, 2, 6, 1, 2},
			k:        3,
			expected: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, DivisibleSumPairs(tt.n, tt.k, tt.ar))
		})
	}
}
