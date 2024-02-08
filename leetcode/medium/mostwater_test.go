package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "test 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "test 2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "test 3",
			height:   []int{4, 3, 2, 1, 4},
			expected: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, MaxArea(tt.height))
		})
	}
}
