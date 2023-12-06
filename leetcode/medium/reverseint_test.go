package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "test 1",
			x:    123,
			want: 321,
		},
		{
			name: "test 2",
			x:    -123,
			want: -321,
		},
		{
			name: "test 3",
			x:    120,
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ReverseInt(tt.x))
		})
	}
}
