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
			input:  []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			input:  []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			input:  []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			input:  []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
		{
			input:  []int{-10, -1, -18, -19},
			target: -19,
			want:   []int{1, 2},
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
