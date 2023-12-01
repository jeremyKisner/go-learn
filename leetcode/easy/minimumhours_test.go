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
		{
			name:              "test 2",
			initialEnergy:     2,
			initialExperience: 4,
			energy:            []int{1},
			experience:        []int{3},
			expected:          0,
		},
		{
			name:              "test 3",
			initialEnergy:     1,
			initialExperience: 1,
			energy:            []int{1, 1, 1, 1},
			experience:        []int{1, 1, 1, 50},
			expected:          51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MinNumberOfHours(tt.initialEnergy, tt.initialExperience, tt.energy, tt.experience))
		})
	}
}
