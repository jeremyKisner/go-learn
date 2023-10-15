package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthday(t *testing.T) {
	tests := []struct {
		name     string
		input    []int32
		d        int32
		m        int32
		expected int32
	}{
		{
			name:     "test 1",
			input:    []int32{1, 2, 1, 3, 2},
			d:        3,
			m:        2,
			expected: 2,
		},
		{
			name:     "test 2",
			input:    []int32{1, 1, 1, 1, 1, 1},
			d:        3,
			m:        2,
			expected: 0,
		},
		{
			name:     "test 3",
			input:    []int32{4},
			d:        4,
			m:        1,
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Birthday(tt.input, tt.d, tt.m)
			assert.Equal(t, tt.expected, b)
		})
	}
}
