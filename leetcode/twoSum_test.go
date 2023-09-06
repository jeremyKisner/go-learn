package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
		{
			[]int{0, 4, 3, 0},
			0,
			[]int{0, 3},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := TwoSum(tt.input, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() got %v, want %v", got, tt.want)
			}
		})
	}
}
