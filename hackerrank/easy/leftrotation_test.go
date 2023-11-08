package hackerrank

import (
	"reflect"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	tests := []struct {
		name     string
		d        int32
		arr      []int32
		expected []int32
	}{
		{
			name:     "test 1",
			d:        2,
			arr:      []int32{1, 2, 3, 4, 5},
			expected: []int32{3, 4, 5, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateLeft(tt.d, tt.arr); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("RotateLeft() = %v, want %v", got, tt.expected)
			}
		})
	}
}
