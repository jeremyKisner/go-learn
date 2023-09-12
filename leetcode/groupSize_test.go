package leetcode

import (
	"reflect"
	"testing"
)

func TestGroupThePeople(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{
			name:  "test group",
			input: []int{3, 3, 3, 3, 3, 1, 3},
			want: [][]int{
				{5},
				{0, 1, 2},
				{3, 4, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupThePeople(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupThePeople() got %v, want %v", got, tt.want)
			}
		})
	}
}
