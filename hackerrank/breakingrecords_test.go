package hackerrank

import (
	"reflect"
	"testing"
)

func TestBreakingRecords(t *testing.T) {
	tests := []struct {
		name   string
		scores []int32
		want   []int32
	}{
		{
			name:   "test 1",
			scores: []int32{10, 5, 20, 20, 4, 5, 2, 25, 1},
			want:   []int32{2, 4},
		},
		{
			name:   "test 2",
			scores: []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42},
			want:   []int32{4, 0},
		},
		{
			name:   "test 3",
			scores: []int32{0, 9, 3, 10, 2, 20},
			want:   []int32{3, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BreakingRecords(tt.scores); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreakingRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
