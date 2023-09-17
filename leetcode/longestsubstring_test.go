package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "test 1",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "test 2",
			input:    "bbbbb",
			expected: 1,
		},
		{
			name:     "test 3",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "test 4",
			input:    " ",
			expected: 1,
		},
		{
			name:     "test 5",
			input:    "au",
			expected: 2,
		},
		{
			name:     "test 5",
			input:    "dvdf",
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.input); got != tt.expected {
				t.Errorf("LengthOfLongestSubstring() got %v, expected %v", got, tt.expected)
			}
		})
	}
}
