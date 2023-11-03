package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			name:     "test 1",
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2,
		},
		{
			name:     "test 2",
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.50000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, findMedianSortedArrays(tt.nums1, tt.nums2))
		})
	}
}
