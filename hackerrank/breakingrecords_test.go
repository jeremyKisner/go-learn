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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BreakingRecords(tt.scores); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreakingRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
