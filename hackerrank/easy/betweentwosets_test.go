package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBetweenTwoSets(t *testing.T) {
	tests := []struct {
		name     string
		a        []int32
		b        []int32
		expected int32
	}{
		{
			name:     "test 1",
			a:        []int32{2, 6},
			b:        []int32{24, 36},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, BetweenTwoSets(tt.a, tt.b))
		})
	}
}
