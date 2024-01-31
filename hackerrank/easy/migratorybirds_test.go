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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, MigratoryBirds(tt.args))
		})
	}
}
