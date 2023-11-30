package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinNumberOfHours(t *testing.T) {
	tests := []struct {
		name              string
		initialEnergy     int
		initialExperience int
		energy            []int
		experience        []int
		expected          int
	}{
		{
			name:              "test 1",
			initialEnergy:     5,
			initialExperience: 3,
			energy:            []int{1, 4, 3, 2},
			experience:        []int{2, 6, 3, 1},
			expected:          8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MinNumberOfHours(tt.initialEnergy, tt.initialExperience, tt.energy, tt.experience))
		})
	}
}
