package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test 1",
			s:    "babad",
			want: "bab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, LongestPalindrome(tt.s))
		})
	}
}
