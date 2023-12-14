package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test 1",
			nums: []int{3, 1, 2, 4},
			want: []int{4, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SortArrayByParity(tt.nums))
		})
	}
}
