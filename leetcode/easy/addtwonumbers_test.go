package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		num1 int
		num2 int
		want int
	}{
		{
			name: "test 1",
			num1: 12,
			num2: 5,
			want: 17,
		},
		{
			name: "test 2",
			num1: -10,
			num2: 4,
			want: -6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, AddTwoNumbers(tt.num1, tt.num2))
		})
	}
}
