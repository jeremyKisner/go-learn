package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternatively(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected string
	}{
		{
			name:     "test 1",
			word1:    "abc",
			word2:    "pqr",
			expected: "apbqcr",
		},
		{
			name:     "test 2",
			word1:    "abc",
			word2:    "defg",
			expected: "adbecfg",
		},
		{
			name:     "test 3",
			word1:    "",
			word2:    "defg",
			expected: "defg",
		},
		{
			name:     "test 4",
			word1:    "abc",
			word2:    "",
			expected: "abc",
		},
		{
			name:     "test 5",
			word1:    "abcd",
			word2:    "pq",
			expected: "apbqcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MergeAlternatively(tt.word1, tt.word2))
		})
	}
}
