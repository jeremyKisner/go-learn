package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumberOfDifferentIntegers(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{
			name: "Test 1",
			word: "a123bc34d8ef34",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, NumberOfDifferentIntegers(tt.word))
		})
	}
}
