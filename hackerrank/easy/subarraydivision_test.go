package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthday(t *testing.T) {
	tests := []struct {
		name         string
		input        []int32
		maxTotal     int32
		numOfSquares int32
		expected     int32
	}{
		{
			name:         "test 1",
			input:        []int32{1, 2, 1, 3, 2},
			maxTotal:     3,
			numOfSquares: 2,
			expected:     2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Birthday(tt.input, tt.maxTotal, tt.numOfSquares)
			assert.Equal(t, tt.expected, b)
		})
	}
}
