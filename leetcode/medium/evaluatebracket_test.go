package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Evaluate(t *testing.T) {
	type args struct {
		s         string
		knowledge [][]string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "test 1",
			args: args{
				s: "(name)is(age)yearsold",
				knowledge: [][]string{
					{"name", "bob"},
					{"age", "two"},
				},
			},
			expected: "bobistwoyearsold",
		},
		{
			name: "test 2",
			args: args{
				s: "hi(name)",
				knowledge: [][]string{
					{"hi(name)"},
				},
			},
			expected: "hi?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Evaluate(tt.args.s, tt.args.knowledge))
		})
	}
}
