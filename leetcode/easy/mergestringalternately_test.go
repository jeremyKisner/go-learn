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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MergeAlternatively(tt.word1, tt.word2))
		})
	}
}
