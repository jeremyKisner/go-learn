package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		limit    int
		expected int64
	}{
		{
			name:     "test 1",
			n:        5,
			limit:    2,
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, DistributeCandies(tt.n, tt.limit))
		})
	}
}
