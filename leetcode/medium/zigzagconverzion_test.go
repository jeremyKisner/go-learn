package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{
			name:    "test 1",
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, Convert(tt.s, tt.numRows))
		})
	}
}
