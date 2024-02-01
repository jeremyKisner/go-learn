package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMigratoryBirds(t *testing.T) {
	tests := []struct {
		name     string
		args     []int32
		expected int32
	}{
		{
			name:     "test 1",
			args:     []int32{1, 1, 2, 2, 3},
			expected: 1,
		},
		{
			name:     "test 2",
			args:     []int32{1, 4, 4, 4, 5, 3},
			expected: 4,
		},
		{
			name:     "test 3",
			args:     []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4},
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, MigratoryBirds(tt.args))
		})
	}
}
