package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingLeaderboard(t *testing.T) {
	tests := []struct {
		name     string
		ranked   []int32
		player   []int32
		expected []int32
	}{
		{
			name:     "test 1",
			ranked:   []int32{100, 90, 90, 80},
			player:   []int32{70, 80, 105},
			expected: []int32{4, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ClimbingLeaderboard(tt.ranked, tt.player)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
